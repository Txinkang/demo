package _defer

import (
	"fmt"
	"testing"
)

/*不捕获异常，是否继续执行*/
func TestPanic1(t *testing.T) {
	defer_call()
	fmt.Println("main正常结束")
}
func defer_call() {
	defer func() { fmt.Println("defer:panic 之前1") }()
	defer func() { fmt.Println("defer:panic 之前2") }()
	panic("异常内容") //触发defer 出栈
	defer func() { fmt.Println("defer:panic 之后,永远执行不到") }()
}

/*捕获异常，是否继续执行*/
func TestPanic2(t *testing.T) {
	defer_call2()
	fmt.Println("main正常结束")
}
func defer_call2() {
	defer func() {
		fmt.Println("defer:panic 之前1，捕获异常")
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	defer func() { fmt.Println("defer:panic 之前2，不捕获异常") }()
	panic("异常内容") //触发defer 出栈
	defer func() { fmt.Println("defer:panic 之后,永远执行不到") }()
}
