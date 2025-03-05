package main

import (
	"html/template"
	"log"
	"os"
)

type person struct {
	Name string
	Age  int
}
type doubt struct {
	person
	Adult bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	fperson := doubt{
		person{
			Name: "Rahim",
			Age:  17,
		},
		false,
	}

	err := tpl.Execute(os.Stdout, fperson)
	if err != nil {
		log.Fatalln(err)
	}
}
