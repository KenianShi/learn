package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name,"name","world","name you called")  //将命令行接收到的值绑定在某个变量上面
																					// 这里还有一个flag.String函数，其实返回一个将接收到的值绑定在次返回的变量上面
}

func main() {
	flag.Parse()		//一般放在main函数的第一行
	fmt.Printf("Hello, %s!\n",name)

}