package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("temp/temp")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var s1, s2, s3 string
	//Fscanf方法将file中的一行按照format的格式给后面的变量赋值
	fmt.Fscanf(file, "%s", &s1)
	fmt.Fscanf(file, "%s", &s2)
	fmt.Fscanf(file, "%s", &s3)

	fmt.Printf("s1:%s  s2:%s s3:%s", s1, s2, s3)

}
