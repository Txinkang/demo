package _defer

import (
	"fmt"
	"testing"
)

func deferFunc() int {
	fmt.Println("defer func called")
	return 0
}
func returnFunc() int {
	fmt.Println("return func called")
	return 0
}
func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}
func TestDeferReturn(t *testing.T) {
	returnAndDefer()
}
