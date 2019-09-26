package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d \n", id, n)
	}
}

func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Millisecond * 1500)
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	var c1, c2 chan int = generator(), generator()
	w := createWorker(0)
	n := 0
	hasValue := false
	tm := time.After(time.Second * 10)
	var value []int
	tick := time.Tick(time.Second)
	for {
		var activeWorkder chan int
		var activeValue int
		if len(value) > 0 {
			activeWorkder = w
			activeValue = value[0]
		}
		if hasValue {
			activeWorkder = w
		}

		select {
		case n = <-c1:
			value = append(value, n)
		case n = <-c2:
			value = append(value, n)
		case activeWorkder <- activeValue:
			value = value[1:]
		case <-time.After(time.Millisecond * 800):
			fmt.Println("Time out 800s")
		case <-tick:
			fmt.Printf("value length:%d\n", len(value))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
