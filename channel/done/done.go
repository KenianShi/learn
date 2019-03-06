package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in   chan int
	done func()
}

func work(i int, w worker) {
	for {
		n := <-w.in
		fmt.Printf("Worker %d received %d \n", i, n)
		w.done()
	}
}

func createWork(i int, wg *sync.WaitGroup) worker {
	w := worker{make(chan int), func() {
		wg.Done()
	}}
	go work(i, w)
	return w
}

func chanDemo() {
	var works [10]worker
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		works[i] = createWork(i, &wg)
	}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		works[i].in <- i + 1000
	}

	for i := 0; i < 10; i++ {
		works[i].in <- i
	}
	wg.Wait()
}

func main() {
	chanDemo()
}
