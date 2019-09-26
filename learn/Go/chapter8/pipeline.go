package main

import (
	"fmt"
	"time"
)

func main() {
	count := make(chan int)
	square := make(chan int)
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			count <- i
			time.Sleep(time.Millisecond * 500)
		}
		close(count)
	}()

	go func() {
		for c := range count {
			sqa := c * c
			square <- sqa
		}
		close(square)
	}()

	go func() {
		for c := range square {
			fmt.Println(c)
		}
		fmt.Println("打印结束")
		done <- true
	}()
	fmt.Println("waiting...")
	<-done

}
