package escape

import (
	"testing"
)

//go:noinline
func foo2(arg_val int) *int {
	var foo_val1 *int = new(int)
	var foo_val2 *int = new(int)
	var foo_val3 *int = new(int)
	var foo_val4 *int = new(int)
	var foo_val5 *int = new(int)

	println(arg_val, foo_val1, foo_val2, foo_val3, foo_val4, foo_val5)

	//将 foo val3 返给 main 兩数
	return foo_val3
}
func TestEscape2(t *testing.T) {
	main_val := foo2(666)
	println(*main_val, main_val)
}
