package test1

import (
	"shikenian/learn/temp/struct1"
	"fmt"
)

func Test1(){
	log := struct1.GetLogShi()
	fmt.Printf("test1:%s \n",log.Name)
}
