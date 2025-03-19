package main

import (
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

func main() {
	var dog hotdog
	var cat hotcat
	http.Handle("/dog", dog)
	http.Handle("/cat", cat)
	http.ListenAndServe(":8080", nil)
}
