package goroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func busi(ch chan bool, i int) {
	fmt.Println("go func ", i, "goroutine count =", runtime.NumGoroutine())
	<-ch
}
func TestLimitGoroutine(t *testing.T) {
	//模拟用户需求业务的数量
	//taskcnt := math.MaxInt64
	taskcnt := 10
	ch := make(chan bool, 3)
	for i := 0; i < taskcnt; i++ {
		ch <- true
		go busi(ch, i)
	}
}
