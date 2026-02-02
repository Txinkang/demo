package goroutine

import (
	"fmt"
	"math"
	"runtime"
	"testing"
)

func busi3(ch chan bool, i int) {
	fmt.Println("go func ", i, "goroutine count =", runtime.NumGoroutine())
	<-ch
	wg.Done()
}
func TestLimitGoroutine3(t *testing.T) {
	task_cnt := math.MaxInt64
	ch := make(chan bool, 3)
	for i := 0; i < task_cnt; i++ {
		wg.Add(1)
		ch <- true
		go busi3(ch, i)
	}
	wg.Wait()
}
