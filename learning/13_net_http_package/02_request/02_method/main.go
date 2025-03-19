// package main
//
// import (
//
//	"html/template"
//	"log"
//	"net/http"
//	"net/url"
//	"os"
//
// )
//
// type hotdog int
//
//	func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//		err := req.ParseForm()
//		if err != nil {
//			log.Println("Error parsing form:", err)
//			http.Error(w, "Unable to process form", http.StatusBadRequest)
//			return
//			// log.Fatalln(err)
//		}
//		data := struct {
//			Method      string
//			URL         *url.URL
//			Submissions url.Values
//		}{
//			req.Method,
//			req.URL,
//			req.Form,
//		}
//		err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
//		if err != nil {
//			log.Println("Error executing template:", err)
//		}
//	}
//
// var tpl *template.Template
//
//	func init() {
//		// tpl = template.Must(template.ParseFiles("tpl.gohtml"))
//		var err error
//		tpl, err = template.ParseFiles("tpl.gohtml")
//		if err != nil {
//			log.Fatalln("Error Parsing templete:", err)
//		}
//	}
//
//	func main() {
//		var d hotdog
//		log.Println("Server Starting on port 8080: 🚀")
//		err := http.ListenAndServe(":8080", d)
//		if err != nil {
//			log.Fatalln("Server error:", err)
//		}
//	}
package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
		http.Error(w, "Form persing error", http.StatusBadRequest)
		return
	}

	data := struct {
		Method      string
		Submissions url.Values
	}{
		req.Method,
		req.Form,
	}
	err = tpl.ExecuteTemplate(w, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln("Failed parse tempete:", err)
	}
}

var tpl *template.Template

func init() {
	var err error
	tpl, err = template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln("Parsing file error:", err)
	}
}

func main() {
	var d hotdog
	log.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", d)
	if err != nil {
		log.Fatalln("Couldn't reach port 8080", err)
	}
}
