package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	xs := []string{"zero", "one", "two", "three", "four", "five"}

	data := struct {
		Words []string
		Name  string
	}{
		xs,
		"Puspo",
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
