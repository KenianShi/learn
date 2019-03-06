package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	r,_ :=http.Get("https://www.cmttracking.io/address/0x3af427d092f9bf934d2127408935c1455170ea8a")
	fmt.Printf("%s \n",r.Body)
	fmt.Println("================================================================")
	b,_ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s ",b)
}
