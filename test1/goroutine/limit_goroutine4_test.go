package goroutine

import (
	"fmt"
	"math"
	"runtime"
	"testing"
)

func busi4(ch chan int) {
	for t := range ch {
		fmt.Println("go func = ", t, "goroutine count =", runtime.NumGoroutine())
		wg.Done()
	}
}
func sendTask(task int, ch chan int) {
	wg.Add(1)
	ch <- task
}
func TestLimitGoroutine4(t *testing.T) {
	ch := make(chan int)
	goCnt := 3
	for i := 0; i < goCnt; i++ {
		go busi4(ch)
	}
	task_cnt := math.MaxInt64
	for t := 0; t < task_cnt; t++ {
		sendTask(t, ch)
	}
	wg.Wait()
}
