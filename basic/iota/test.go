package main

import (
	"fmt"
)

func p(a ...string){
	fmt.Println(a)
}

func main() {
	s := []string{"XXX","YYY"}
	// p("sdfsd",s...)			//不能把slice切片和元素混合在一起
	s = append(s, "sdfsd")
	p(s...)

	// 拓展
	arr := []string{"AAA","BBB","CCC","DDD"}
	x := "XXX"
	after := append([]string{},arr[2:]...)		//这里需要使用这种方法吧slice的值复制出来
	arr = append(append(arr[:2], x), after...)
	fmt.Println(arr)
}
