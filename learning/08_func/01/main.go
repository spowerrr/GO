// package main
//
// import (
//
//	"log"
//	"os"
//	"strings"
//	"text/template"
//
// )
//
// var tpl *template.Template
//
//	type sage struct {
//		Name  string
//		Motto string
//	}
//
//	type car struct {
//		Manufacturer string
//		Model        string
//		Doors        int
//	}
//
// // create a func map to register functions
// // "uc" is what the func will be called in the temple
// // "uc" is the toUppper func from package strings
// // "ft"is a func i declared
// // "ft" slices  a string returning the firest
//
//	var fm = template.FuncMap{
//		"uc": strings.ToUpper,
//		"ft": firstThree, // firstThree is a function
//	}
//
// // not needed when refactoring
// // type items struct {
// // 	Wisdom    []sage
// // 	Transport []car
// // }
//
// //example how to use temple New
// // func main(){
// // 	tpl:= template.Must(template.New("this").Parse("Hi there"))
// // 	tpl.ExecuteTemplate(os.Stdout, "this", nil) }
// //--> for this output: Hi there
//
//	func init() {
//		tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
//	}
//
//	func firstThree(s string) string {
//		s = strings.TrimSpace(s)
//		if len(s) >= 3 {
//			s = s[:3] //-->first there characters
//		}
//		return s
//	}
//
//	func main() {
//		a := sage{
//			Name:  "Buddha",
//			Motto: "The belief of no beliefs",
//		}
//		b := sage{
//			Name:  "Gandhi",
//			Motto: "Be the change",
//		}
//		c := sage{
//			Name:  "Martin Luthar King",
//			Motto: "Hatred never ceases with hatred but with love alone is healed.",
//		}
//		f := car{
//			Manufacturer: "Ford",
//			Model:        "F150",
//			Doors:        2,
//		}
//		g := car{
//			Manufacturer: "Toyota",
//			Model:        "Corolla",
//			Doors:        3,
//		}
//		sages := []sage{a, b, c}
//		cars := []car{f, g}
//		// refactor
//		items := struct {
//			Wisdom    []sage
//			Transport []car
//		}{
//			sages,
//			cars,
//		}
//		err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", items)
//		if err != nil {
//			log.Fatalln(err)
//		}
//	}
package main

import (
	"log"
	"os"
	"strings"
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

// create a FuncMap to register functions.
// "uc" is what the func will be called in the template
// "uc" is the ToUpper func from package strings
// "ft" is a func I declared
// "ft" slices a string, returning the first three characters
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("tpl").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {
	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
