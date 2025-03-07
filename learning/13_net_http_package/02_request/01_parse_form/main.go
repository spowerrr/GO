package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	} // Check if the form was submitted via POST method
	if req.Method == http.MethodPost {
		// Render the template with the parsed form data
		tpl.ExecuteTemplate(w, "tpl.gohtml", req.Form)
	} else {
		// Serve the form page if not a POST request
		tpl.ExecuteTemplate(w, "tpl.gohtml", nil)
	}
	/* tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", req.Form) */
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
