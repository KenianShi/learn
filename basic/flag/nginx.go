package main

import (
	"flag"
	"fmt"
)

var A bool

var (
	h    bool
	v, V bool		// 导出的（大写的）变量应该有自己的变量定义，而不适合在一个var里面进行定义
	t, T *bool
	q    *bool
	s    *string
	p    *string
)

func init() {
	flag.BoolVar(&h, "h", false, "this is help")
	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.BoolVar(&V, "V", false, "show version and configure option then exit")
	t = flag.Bool("t", false, "test configuration and exit")
	T = flag.Bool("T", false, "test configuration,dump it and exit")
	q = flag.Bool("q", false, "suppress non-error message during configuration")
	s = flag.String("s","","send `signal` to a master process:stop, quitr, reopen, reload")
	p = flag.String("p","/usr/local/nginx/","set `prefix` path")
}

func main() {
	flag.Parse()
	fmt.Println("parsed? = ",flag.Parsed())

	fmt.Println(*s)



}
