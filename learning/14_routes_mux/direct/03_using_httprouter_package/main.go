package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templetes/*"))
}

// user,blogread,bloagWrite we need those val
func user(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "USER: %s!\n", p.ByName("name"))
}

func blogWriter(w http.ResponseWriter, r *http.Request, writer httprouter.Params) {
	fmt.Fprintf(w, "READ CATEGORY, %s!\n", writer.ByName("category"))
	fmt.Fprintf(w, "READ ARTICLE, %s!\n", writer.ByName("article"))
}

func blogReader(w http.ResponseWriter, r *http.Request, reader httprouter.Params) {
	fmt.Fprintf(w, "READ CATEGORY, %s!\n", reader.ByName("category"))
	fmt.Fprintf(w, "READ ARTICLE, %s!\n", reader.ByName("article"))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func applyProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}

// func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { //example user
//
//		fmt.Fprintf(w, "Hello, %s!\n", ps.ByName("name"))
//	}

// function to handle all the errors -> for all the functions
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/user/:name", user)
	router.GET("/about", about)
	router.GET("/apply", apply)
	router.POST("/apply", applyProcess)
	router.GET("/contact", contact)
	router.POST("/blogWriter/:category/:article", blogWriter)
	router.GET("/blogReader/:category/:article", blogReader)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
