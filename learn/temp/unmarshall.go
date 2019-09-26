package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	A 	string  	`json:A`
	B 	int	`json:B`
}

func main() {
	var z A
 	s := `{"A":"A","B":"0x12"}`
	err := json.Unmarshal([]byte(s),&z)
	if err != nil{
		fmt.Printf("%T,%v \n",z,z)
	}else {

		fmt.Println(err)
	}


}
