package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//t1 := time.Now().UnixNano()  //格林威治时间，单位纳秒    1秒=1000毫秒=1000 000 微秒=1000 000 000纳秒
	//t2 := time.Now().Unix()		//格林威治时间，单位秒
	//fmt.Printf("t1: %d \n",t1)
	//fmt.Printf("t2: %d",t2)
	var a int64 = 2
	rand.Seed(a)
	x1 := rand.Int()
	x2 := rand.Int()
	fmt.Println(x1)
	fmt.Println(x2)
	//1543039099823358511
	//2444694468985893231

}
