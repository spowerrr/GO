package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type country struct {
	Name    string
	Capital string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	cr := country{
		Name:    "Bangladesh",
		Capital: "Dhaka",
	}
	err := tpl.Execute(os.Stdout, cr)
	if err != nil {
		log.Default().Fatalln(err)
	}
}
