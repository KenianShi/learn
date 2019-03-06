package main

import "fmt"

func f(i int) func() int {
	return func() int {
		i++
		return i
	}
}

func main() {
	f1 := f(10)
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	f2 := f(20)
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
	f2 = f(20)
	fmt.Println("f2():", f2())

}
