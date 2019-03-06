package main

import (
	"fmt"
	"regexp"
)

const text = `this is shikenian@163.com@com
		tomshi@qq.com
		goodmorning@139.com
		owen@abc.com.cn
`

func main() {
	reg := regexp.MustCompile(`([A-Za-z0-9]+)@([a-zA-Z0-9]+).([A-Za-z0-9.]+)`)
	fmt.Println(reg.FindString(text))
	fmt.Println("========================")
	stringArray := reg.FindAllString(text, -1)
	for _, v := range stringArray {
		fmt.Println(v)
	}
	fmt.Println("********************************************")
	strings := reg.FindAllStringSubmatch(text, -1)
	for _, v := range strings {
		fmt.Println(v)
	}

}
