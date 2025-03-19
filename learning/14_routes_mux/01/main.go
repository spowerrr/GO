package main

import (
	"fmt"
	"io"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "HotDOOOG")
}

type hotcat int

func (h hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "HotCAAAAT")
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "This is my home page")
}

func main() {
	var dog hotdog
	var cat hotcat
	// routing using mux
	mux := http.NewServeMux()
	// Fix: Use http.HandlerFunc for handling "/"
	mux.Handle("/", http.HandlerFunc(homeHandler))
	mux.Handle("/cat", cat)
	mux.Handle("/dog", dog)
	http.ListenAndServe(":8080", mux)
}
