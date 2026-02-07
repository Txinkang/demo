package mem

import (
	"fmt"
	"testing"
)

func TestZero(t *testing.T) {
	var (
		//0 内存对象
		a struct{}
		b [0]int
		//100个0内存 struct{}
		c [100]struct{}
		//100个0内存 struct{},make 申请形式
		d = make([]struct{}, 100)
	)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", &c[50])
	fmt.Printf("%p\n", &(d[50]))
}
