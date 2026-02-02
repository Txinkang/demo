package escape

import "testing"

//go:noinline
func foo(arg_val int) *int {
	var foo_val1 int = 11
	var foo_val2 int = 12
	var foo_val3 int = 13
	var foo_val4 int = 14
	var foo_val5 int = 15
	//此处循环为了防止go编译器将foo优化成inline(内联函数)
	//如果是内联函数,则 main调用foo将是原地展开,所以foo_val 1-5相当于main作用域的变量
	//即使foo_val3发生逃逸,地址与其他也是连续的
	for i := 0; i < 5; i++ {
		println(&arg_val, &foo_val1, &foo_val2, &foo_val3, &foo_val4, &foo_val5)
	}
	//将 foo val3 返给 main 兩数
	return &foo_val3
}
func TestEscape(t *testing.T) {
	main_val := foo(666)
	println(*main_val, main_val)
}
