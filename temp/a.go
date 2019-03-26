package main

import (
	"fmt"
	"encoding/json"
)

func main(){

	type student struct {
		Name string
		Age int
	}

	a:=&student{
		Name:"张三",
		Age:12,
	}
	aBytes,err:=json.Marshal(a)
	if err != nil {
		panic("++++++")
	}
	fmt.Println("======",string(aBytes))
}