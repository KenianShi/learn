package main

import (
	"os"
	"os/signal"
	"fmt"
)

func main() {
	sigs := make(chan os.Signal,1)
	done := make(chan bool,1)
	signal.Notify(sigs)		//会将后面填写的信号（syscall.SIGINT,syscall.SIGTERM）传递给第一个参数
	go func(){
		sig := <- sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	fmt.Println("awaiting signal")
	<- done
	fmt.Println("exiting")



}
