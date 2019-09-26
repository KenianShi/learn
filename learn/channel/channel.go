package main

import (
	"fmt"
	"time"
)

func createChannel() chan int {
	c := make(chan int)
	go func() {
		for {
			fmt.Println("Worker %d received %c \n")
		}
	}()
	return c

}

func work(id int, a chan int) {
	for {
		n := <-a
		fmt.Printf("Worker %d received %c \n", id, n)
	}
}

func worker(c chan int) {
	for i := range c {
		fmt.Println(i)
	}
}

func bufferChannel() {
	channel := make(chan int, 3)
	go worker(channel)
	channel <- 'a'
	channel <- 'b'
	channel <- 'H'
	channel <- 'X'
	close(channel)
	time.Sleep(time.Second)

}

func chanDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go work(i, channels[i])
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Second)

}

func main() {
	//chanDemo()
	bufferChannel()
}
