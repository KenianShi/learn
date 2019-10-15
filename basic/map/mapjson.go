package main

import (
	"encoding/json"
	"fmt"
)

type wallet struct {
	key 	map[string]interface{}
}

type A struct {

}

func main() {
	a := A{}
	w1 := wallet{key: map[string]interface{}{"bob":a}}
	bz,err := json.Marshal(w1)
	if err != nil {
		panic(err)
	}
	w2 := wallet{}
	err  = json.Unmarshal(bz,&w2)
	if err != nil {
		panic(err)
	}
	fmt.Println("ok")
}