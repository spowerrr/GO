package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template //It creates a nil pointer at first — which means the variable doesn't point to any template yet.
//Later in the init() function, this pointer gets assigned to the templates using

//var tpl template.Template -->can't use because It would create a new empty template instance.

func init() {
	tpl = template.Must(template.ParseGlob("templete/*")) //func Must takes *template,err and returns * template
}
func main() {

	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}

//Step-by-step flow:

//1.tpl is declared as a nil pointer.
//2.Inside init(), template.ParseGlob() reads all template files from "templete/*".
//3.ParseGlob() returns a pointer to *template.Template.
//4.template.Must() makes sure there are no errors during parsing.
//5.Now tpl holds a reference (pointer) to all parsed templates.
