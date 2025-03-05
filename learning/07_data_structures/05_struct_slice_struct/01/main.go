package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}
type car struct {
	Manufacturer string
	Model        string
	Doors        int
}
type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	a := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}
	b := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}
	c := sage{
		Name:  "Martin Luthar King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}
	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}
	g := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        3,
	}
	sages := []sage{a, b, c}
	cars := []car{f, g}
	items := items{
		Wisdom:    sages,
		Transport: cars,
	}
	err := tpl.Execute(os.Stdout, items)
	if err != nil {
		log.Fatalln(err)
	}
}
