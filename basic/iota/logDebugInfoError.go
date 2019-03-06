package main

import "fmt"

type level int

var arr = []string{"debug","info","warn","error"}

func creatHandle(i level){
	for i<4{
		fmt.Println(arr[i])
		i++
	}
}

func main() {
	creatHandle(1)
}
