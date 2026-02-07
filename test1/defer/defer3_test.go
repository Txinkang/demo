package _defer

import (
	"fmt"
	"testing"
)

/*测试初始化是否为0*/
func DeferFunc1(i int) (t int) {
	fmt.Println("t= ", t)
	return 2
}
func TestValue(t *testing.T) {
	DeferFunc1(10)
}

/*测试defer是否可以修改返回值*/
func returnButDefer() (t int) {
	defer func() {
		t = t * 10
	}()
	return 1
}
func TestChangeValue(t *testing.T) {
	fmt.Println(returnButDefer())
}
