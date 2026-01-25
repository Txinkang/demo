package mem

import (
	"errors"
	"fmt"
	"sync"
)

// Pool 内存管理池类型
type Pool map[int]*Buf

// BufPool Buf 内存池
type BufPool struct {
	//所有 buffer 的一个 map 集合句柄
	Pool     Pool
	PoolLock sync.RWMutex

	//总 buffer 池的内存大小 单位为 KB
	TotalMem uint64
}

// 单例对象
var bufPoolInstance *BufPool
var once sync.Once

// MemPool 获取 BufPool 对象(单例模式)
func MemPool() *BufPool {
	once.Do(func() {
		bufPoolInstance = new(BufPool)
		bufPoolInstance.Pool = make(map[int]*Buf)
		bufPoolInstance.TotalMem = 0
		//bufPoolInstance.prev = nil
		bufPoolInstance.initPool()
	})
	return bufPoolInstance
}

const (
	m4K   int = 4096
	m16K  int = 16384
	m64K  int = 65535
	m256K int = 262144
	m1M   int = 1048576
	m4M   int = 4194304
	m8M   int = 8388608
)

/*
初始化内存池 主要是预先开辟一定量的空间这里 BufPool 是一个 hash, 每个 key 都是不同空间容量对应的 value 是一个 Buf 集合的链表

	BufPool -->[m4K]-- Buf-Buf-Buf-Buf...(BufList)
	[m16K]-- Buf-Buf-Buf- Buf...(BufList)
	[m64K] --Buf-Buf-Buf- Buf...(BufList)
	[m256K]-- Buf-Buf-Buf- Buf...(BufList)
	[m1M]-- Buf-Buf-Buf- Buf...(BufList)
	[m4M]-- Buf-Buf-Buf-Buf...(BufList)
	[m8M]-- Buf-Buf-Buf-Buf...(BufList)
*/
func (bp *BufPool) initPool() {
	//---->开辟 4KB buf 内存池
	//4KB 的 Buf 预先开辟 5000个,约 20MB 供开发者使用
	bp.makeBufList(m4K, 5000)
	//---->开辟 16KB buf 内存池
	//16KB 的 Buf 预先开辟 1000 个,约 16MB 供开发者使用
	bp.makeBufList(m16K, 1000)
	//---->开辟 64KB buf 内存池
	//64KB的Buf预先开辟500个, 约 32MB 供开发者使用
	bp.makeBufList(m64K, 500)
	//---->开辟 256KB buf 内存池
	//256KB的 Buf 预先开辟 200个,约 50MB 供开发者使用
	bp.makeBufList(m256K, 200)
	//---->开辟 1MB buf 内存池
	//1MB 的 Buf 预先开辟 50个,约 50MB 供开发者使用
	bp.makeBufList(m1M, 50)
	//---->开辟 4MB buf 内存池
	//4MB 的 Buf 预先开辟 20个，约 80MB 供开发者使用
	bp.makeBufList(m4M, 20)
	//---->开辟 8MB buf 内存池
	//8MB的 io buf 预先开辟 10个,约 80MB 供开发者使用
	bp.makeBufList(m8M, 10)
}

func (bp *BufPool) makeBufList(cap int, num int) {
	bp.Pool[cap] = NewBuf(cap)
	var prev *Buf
	prev = bp.Pool[cap]
	for i := 1; i < num; i++ {
		prev.Next = NewBuf(cap)
		prev = prev.Next
	}
	bp.TotalMem += (uint64(cap) / 1024) * uint64(num)
}

//总内存池最大限制 单位是 KB,所以目前的限制是 5GB

const EXTRA_MEM_LIMIT int = 5 * 1024 * 1024

// Alloc 开辟一个 Buf
func (bp *BufPool) Alloc(N int) (*Buf, error) {
	//1 找到N最接近哪个 hash 组
	var index int
	if N <= m4K {
		index = m4K
	} else if N <= m16K {
		index = m16K
	} else if N <= m64K {
		index = m64K
	} else if N <= m256K {
		index = m256K
	} else if N <= m1M {
		index = m1M
	} else if N <= m4M {
		index = m4M
	} else if N <= m8M {
		index = m8M
	} else {
		return nil, errors.New("Alloc size Too Large!")
	}
	//2如果该组已经没有，则需要额外申请，所以需要加锁保护
	bp.PoolLock.Lock()
	if bp.Pool[index] == nil {
		if (bp.TotalMem + uint64(index/1024)) >= uint64(EXTRA_MEM_LIMIT) {
			errStr := fmt.Sprintf("already use too many memory! \n")
			return nil, errors.New(errStr)
		}
		newBuf := NewBuf(index)
		bp.TotalMem += uint64(index / 1024)
		bp.PoolLock.Unlock()
		fmt.Printf("Alloc Mem Size: %d KB\n", newBuf.Capacity/1024)
		return newBuf, nil
	}
	//3 如果该组有 Buf 内存存在,则得到一个 Buf 并返回,并且从 pool 中移除该内存块
	targetBuf := bp.Pool[index]
	bp.Pool[index] = targetBuf.Next
	bp.TotalMem -= uint64(index / 1024)
	bp.PoolLock.Unlock()
	targetBuf.Next = nil
	fmt.Printf("Alloc Mem Size: %d KB\n", targetBuf.Capacity/1024)
	return targetBuf, nil
}

// Revert 当 alloc 之后,当前 Buf 被使用完,需要重置这个 Buf,并且需要将该 buf 放回 pool 中
func (bp *BufPool) Revert(buf *Buf) error {
	//每个 buf 的容量都是固定的，在 hash的 key 中取值
	index := buf.Capacity
	//重置 buf 中的内置位置指针
	buf.Clear()
	bp.PoolLock.Lock()
	//找到对应的 hash组 buf 首节点地址
	if _, ok := bp.Pool[index]; !ok {
		errStr := fmt.Sprintf("Index %d not in BufPoll!\n", index)
		return errors.New(errStr)
	}
	//将 buffer 插回链表头部
	buf.Next = bp.Pool[index]
	bp.Pool[index] = buf
	bp.TotalMem += uint64(index / 1024)
	bp.PoolLock.Unlock()
	fmt.Printf("Revert Mem Size: %d KB\n", index/1024)
	return nil
}
