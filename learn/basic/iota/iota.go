package main

import "fmt"

type LogLevel int

const(
	LogLevelCrit LogLevel = iota
	LogLevelError
	LogLevelWarn
	LogLevelInfo
	LogLevelDebug
)

func main() {
	fmt.Println(LogLevelCrit)
	fmt.Println(LogLevelDebug)
	a := LogLevelDebug + 1
	fmt.Println(a)
	fmt.Printf("%T:%v",a,a)

}