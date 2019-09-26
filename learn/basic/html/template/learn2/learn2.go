package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	t,err := template.ParseFiles("./basic/html/template/learn2/tpl.html")
	if err != nil {
		log.Fatal(err)
	}
	data := struct {
		Title	string
	}{"shikenian"}
	err = t.Execute(os.Stdout,data)
	if err != nil {
		log.Fatal(err)
	}

}
