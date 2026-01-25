package mem

import "fmt"

// 应用层的 buffer 数据
type ZBuf struct {
	b *Buf
}

// 清空当前的 ZBuf
func (zb *ZBuf) Clear() {
	if zb.b != nil {
		//将 Buf 重新放回 buf_pool 中
		MemPool().Revert(zb.b)
		zb.b = nil
	}
}

// 弹出已使用的有效长度
func (zb *ZBuf) Pop(len int) {
	if zb.b == nil || len > zb.b.Length() {
		return
	}
	zb.b.Pop(len)
	//当此时 Buf 的可用长度已经为0时，将 Buf 重新放回 BufPool 中
	if zb.b.Length() == 0 {
		MemPool().Revert(zb.b)
		zb.b = nil
	}
}

// 获取 Buf 中的数据
func (zb *ZBuf) Data() []byte {
	if zb.b == nil {
		return nil
	}
	return zb.b.GetBytes()
}

// 重置缓冲区
func (zb *ZBuf) Adjust() {
	if zb.b != nil {
		zb.b.Adjust()
	}
}

// 将数据读取到 Buf 中
func (zb *ZBuf) Read(src []byte) (err error) {
	if zb.b == nil {
		zb.b, err = MemPool().Alloc(len(src))
		if err != nil {
			fmt.Println("pool Alloc Error ", err)
		}
	} else {
		if zb.b.Head() != 0 {
			return nil
		}
		if zb.b.Capacity-zb.b.Length() < len(src) {
			//不够存放，重新从内存池申请
			newBuf, err := MemPool().Alloc(len(src) + zb.b.Length())
			if err != nil {
				return nil
			}
			//将之前的 Buf 复制到新申请的 Buf 中
			newBuf.Copy(zb.b)
			//将之前的 Buf 回收到内存池中
			MemPool().Revert(zb.b)
			//新申请的 Buf 成为当前的 zBuf
			zb.b = newBuf
		}
	}
	//将内容写进 ZBuf 缓冲中
	zb.b.SetBytes(src)
	return nil
}
