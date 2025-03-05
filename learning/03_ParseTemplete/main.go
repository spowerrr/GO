package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("one.gmao")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	tpl1, err := template.ParseFiles("two.gmao", "three.gmao")
	if err != nil {
		log.Fatalln("Couldn't parse:", err)
	}
	err = tpl1.ExecuteTemplate(os.Stdout, "three.gmao", nil)
	if err != nil {
		log.Fatalln("Couldn't execute:", err)
	}
	err = tpl1.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	if err != nil {
		log.Fatalln("Couldn't execute:", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	if err != nil {
		log.Fatalln("Couldn't execute:", err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
