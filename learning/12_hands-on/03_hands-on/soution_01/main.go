package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type hotel struct {
	Name, Address, City, Zip, Region string
}
type hotels []hotel

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	a := hotels{
		hotel{
			Name:    "Hotel California",
			Address: "42 Sunset Boulevard",
			City:    "Los Angelos",
			Zip:     "95612",
			Region:  "southern",
		}, hotel{
			Name:    "H",
			Address: "4",
			City:    "L",
			Zip:     "95612",
			Region:  "southern",
		},
	}
	err := tpl.Execute(os.Stdout, a)
	if err != nil {
		log.Fatalln(err)
	}
}
