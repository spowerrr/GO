package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	file, err := os.Create("index.html")
	if err != nil {
		fmt.Println("Couldn't create index.html", err)
	}
	defer file.Close()

	err = tpl.Execute(file, 50)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Written code successfully")
	}

}
