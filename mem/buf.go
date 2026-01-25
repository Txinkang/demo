package mem

import "C"
import (
	"fmt"
	"unsafe"
	"zmem/c"
)

type Buf struct {
	//如果存在多个 buffer,则采用链表的形式连接起来
	Next *Buf
	//当前 buffer 的缓存容量大小，也是整个内存的尾地址索引
	Capacity int
	//当前 buffer 的有效数据长度，也是尾地址索引
	length int
	//未处理数据（有效数据）的头部位置索引
	head int
	//当前 buf 所保存的数据地址
	data unsafe.Pointer
}

// NewBuf 构造，创建一个 Buf 对象
func NewBuf(size int) *Buf {
	return &Buf{
		Capacity: size,
		length:   0,
		head:     0,
		Next:     nil,
		data:     c.Malloc(size),
	}
}

// SetBytes 给一个 Buf 填充[ ]Byte 数据
func (b *Buf) SetBytes(src []byte) {
	c.Memcpy(unsafe.Pointer(uintptr(b.data)+uintptr(b.head)), src, len(src))
	b.length += len(src)
}

// GetBytes 获取一个 Buf 的数据,以[]Byte 形式展现
func (b *Buf) GetBytes() []byte {
	data := C.GoBytes(unsafe.Pointer(uintptr(b.data)+uintptr(b.head)), C.int(b.length))
	return data
}

// Copy 将其他 Buf 对象数据复制到自己中
func (b *Buf) Copy(other *Buf) {
	c.Memcpy(b.data, other.GetBytes(), other.length)
	b.head = 0
	b.length = other.length
}

// Pop 处理长度为 len 的数据,移动 head 和修正 length
func (b *Buf) Pop(len int) {
	if b.data == nil {
		fmt.Printf("pop data is nil")
		return
	}
	if len > b.length {
		fmt.Printf("pop len > length")
		return
	}
	b.length -= len
	b.head += len
}

// Adjust 将已经处理过的数据清空，将未处理的数据提至数据首地址
func (b *Buf) Adjust() {
	if b.head != 0 {
		if b.length != 0 {
			c.Memmove(b.data, unsafe.Pointer(uintptr(b.data)+uintptr(b.head)), b.length)
			b.head = 0
		}
	}
}

// 清空数据
func (b *Buf) Clear() {
	b.length = 0
	b.head = 0
}

func (b *Buf) Head() int {
	return b.head
}
func (b *Buf) Length() int {
	return b.length
}
