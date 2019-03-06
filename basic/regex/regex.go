package main

import (
	"fmt"
	"regexp"
)

const text = "My email is shikenian@qq.com"
const test = "ABCDabcd"
const A = ` shikenian@qq.com
			1035868500@qq.com
			123456789@163.com
			shikenian@google.com
`

func main() {
	reg := regexp.MustCompile(`([a-zA-Z0-9]+)@([A-z0-9]+).([A-z0-9]+)`)
	s := reg.FindString(test)
	fmt.Printf("s: %s\n", s)
	reg1 := regexp.MustCompile(`[A-z]+`)
	s1 := reg1.FindString(test)
	fmt.Printf("s1:%s\n", s1)
	fmt.Println("**************************")
	reg2 := reg.FindAllStringSubmatch(A, -1)
	for i := range reg2 {
		fmt.Println(reg2[i])
	}
}
