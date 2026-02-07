package mem

import (
	"log"
	"time"
)

func test() {
	//slice会动态扩容,用slice来做堆内存申请
	container := make([]int, 8)
	log.Println(" ===> loop begin.")
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i)
		log.Println(" ===> loop end.")
	}
}
func CheckMem() {
	log.Println("Start.")
	test()
	log.Println("force gc.") //强制调用runtime.GC()
	log.Println("Done.")
	//睡眠,保持程序不退出
	time.Sleep(3600 * time.Second)
}
