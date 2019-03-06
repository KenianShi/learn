package main

import (
	"time"
	"fmt"
)

func main() {
	done := make(chan int,1)

	go func(){
		time.Sleep(time.Second*5)
		close(done)
	}()

	fmt.Println("prepare end...")
	a := <- done								//从一个关闭的通道里取数据可以，可以得到一个默认值
	//done <- 5								//	往一个关闭的通道里面写数据，会panic
	fmt.Println(a)
	fmt.Println("END")

}
