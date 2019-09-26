package main

import (
	"log"
	"html/template"
	"os"
)

func main() {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<Title>{{.Title}}</Title>
	</head>
	<body>
		{{range .Items}}
		<div>
			{{. }}
		</div>
		{{else}}
		<div>
			<strong>no rows</strong>
		</div>
		{{end}}
	</body>
</html>`

	check := func(err error){
		if err != nil {
			log.Fatal(err)
		}
	}

	t,err := template.New("webpage").Parse(tpl)
	check(err)
	data := struct {
		Title	string
		Items	[]string
	}{
		Title:"shikenian",
		Items:[]string{"one","two","three","four","five"},
	}
	err = t.Execute(os.Stdout,data)
	check(err)
}