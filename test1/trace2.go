package main

import (
	"fmt"
	"time"
)

func trace() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
	}
	fmt.Println("Hello World")

	//下面这部分粘贴到main文件用，只要能执行就行。
	//创建 trace 文件
	//f, err := os.Create("trace.out")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	////启动 trace goroutine
	//err = trace.Start(f)
	//if err != nil {
	//	panic(err)
	//}
	//defer trace.Stop()
	//fmt.Println("Hello World")
}
