package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := [][]byte{[]byte("wujun"),[]byte("xiaoming")}
	fmt.Printf("%s \n",a)
	b := bytes.Join(a,[]byte("-"))			//将两个字节以-分隔号分隔开，最后返回一个字节
	fmt.Printf("%s: \n",b)

}
