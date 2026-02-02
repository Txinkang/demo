package goroutine

import (
	"testing"
	"time"
)

func TestGoroutineSize(t *testing.T) {
	for i := 0; i < 200000; i++ {
		go func() {
			time.Sleep(5 * time.Second)
		}()
		time.Sleep(10 * time.Second)
	}
}
