package main

import (
	"fmt"
	"time"
)

func count(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
		time.Sleep(time.Millisecond * 500)
	}
	close(out)
}

func square(in <-chan int, out chan<- int) {
	for c := range in {
		out <- c * c
	}
	close(out)
}

func printer(in <-chan int, done chan<- bool) {
	for c := range in {
		fmt.Println(c)
	}
	fmt.Println("即将done...")
	done <- true

}

func main() {
	counts := make(chan int)
	squares := make(chan int)
	done := make(chan bool)
	go count(counts)
	go square(counts, squares)
	go printer(squares, done)
	<-done
}
