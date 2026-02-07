package mem

import (
	"fmt"
	"sync"
	"testing"
)

/*
测试new出来的值是不是默认0
*/
type user struct {
	lock sync.Mutex
	name string
	age  int
}

func TestMem(t *testing.T) {
	u := new(user)
	u.lock.Lock()
	u.name = "jack"
	u.lock.Unlock()
	fmt.Println(u)
}
