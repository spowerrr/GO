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
	capital := map[string]string{
		"India":      "Delhi",
		"Bangladesh": "Dhaka",
		"Nepal":      "Kathmandu",
	}
	err := tpl.Execute(os.Stdout, capital)
	if err != nil {
		log.Fatalln(err)
	}
}
