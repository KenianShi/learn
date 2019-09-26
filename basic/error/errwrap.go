package main

import (
	"fmt"
	"github.com/pkg/errors"
)

//github.com/pkg/errors的使用
// errors.wrap最后返回err，其形式是msg:err
func main() {
	a := 1
	err1 := fmt.Errorf("my err1= %v",a)
	fmt.Println(err1)
	err2 := errors.Wrap(err1,"my err2 include err1")
	fmt.Println(err2)
}
