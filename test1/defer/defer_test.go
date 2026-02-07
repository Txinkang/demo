package _defer

import (
	"fmt"
	"testing"
)

/*测试执行顺序*/
func TestDefer(t *testing.T) {
	defer func1()
	defer func2()
	defer func3()
}
func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}
func func3() {
	fmt.Println("C")
}
