package goroutine

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"testing"
)

var wg = sync.WaitGroup{}

func busi2(i int) {
	fmt.Println("go func ", i, "goroutine count =", runtime.NumGoroutine())
	wg.Done()
}
func TestLimitGoroutine2(t *testing.T) {
	task_cnt := math.MaxInt64

	for i := 0; i < task_cnt; i++ {
		wg.Add(1)
		go busi2(i)
	}
	wg.Wait()
}
