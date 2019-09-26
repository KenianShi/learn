package main

import (
	"fmt"
	"time"
)

var g chan int
var quit chan chan bool

func main() {
	g = make(chan int)
	quit = make(chan chan bool)
	go C()
	for i := 0; i <= 5; i++ {
		g <- i
		time.Sleep(time.Millisecond * 500)
	}
	waitB := make(chan bool)
	quit <- waitB
	<- waitB
	fmt.Println("Main quit")
}

func C(){
	for {
		select{
		case i := <- g:
			fmt.Println(i + 1)
		case c := <- quit :
			c <- true
			fmt.Println("B quit")
			return
		}
	}
}

