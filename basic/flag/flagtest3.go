package main

import (
	"flag"
	"fmt"
)

//对于flag绑定的参数，要用-或者--才能被系统识别
func main() {
	//flag.String 返回一个指针，而flag.StringVar则需要先定义一个变量，在将解析的值绑到给定的变量
	//flag.String在其内部生成了一个新的string，作为变量，最后也调用的flag.StringVar()
	wordPtr := flag.String("word","fool","a string")
	numPtr := flag.Int("number",21,"an int")
	boolPtr := flag.Bool("fork",false,"a bool")
	var svar string
	flag.StringVar(&svar,"svar","bar","a string var")


	flag.Parse()
	fmt.Println("word",*wordPtr)
	fmt.Println("number",*numPtr)
	fmt.Println("fork",*boolPtr)
	fmt.Println("svar",svar)
	fmt.Println("tail",flag.Args())


}
