package main

import (
	"net/http"
	"fmt"
	"log"
)

func index(w http.ResponseWriter,res *http.Request){
	fmt.Fprintf(w, " Hello golang http!")
}

func main() {
	http.HandleFunc("/",index)
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}

}