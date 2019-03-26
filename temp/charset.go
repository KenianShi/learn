package main

import "fmt"

const charset1 = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

func main() {
	data := []byte("1cde9j30h8teunj2fcz422y8lr0e2xu5q")
	fmt.Println(data)
	for _,b := range data{
		if int(b) >= len(charset1) {
			//fmt.Println(b)
			fmt.Print(int(b),"-")
			fmt.Println(string(b))
		}else{
		fmt.Println("aaa:",string(charset1[b]))
	}}



}
