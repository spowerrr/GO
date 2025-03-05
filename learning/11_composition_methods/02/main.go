package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type course struct {
	Number, Name, Unit string
}
type semester struct {
	Term    string
	Courses []course
}
type year struct {
	Fall, Spring, Summer semester
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	one := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{
					"CSE 1110",
					"Introduction to Computer Systems",
					"1",
				},
				{
					"CSE 1111",
					"Structured Programming Language",
					"3",
				},
				{
					"CSE 1112",
					"Structured Programming Language Laboratory",
					"1",
				},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{
					"CSE 1115",
					"Object Oriented Programming",
					"3",
				},
				{
					"CSE 1116",
					"Object Oriented Programming Laboratory",
					"1",
				},
				{
					"CSE 2118",
					"Advanced Object Oriented Programming Laboratory",
					"1",
				},
			},
		},
	}
	err := tpl.Execute(os.Stdout, one)
	if err != nil {
		log.Fatalln(err)
	}
}
