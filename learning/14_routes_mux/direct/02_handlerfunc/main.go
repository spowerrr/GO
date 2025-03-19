package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "DOGGGGG")
}

func h(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hOG")
}

func main() {
	http.Handle("/dog", http.HandlerFunc(d))
	http.Handle("/hog", http.HandlerFunc(h))
	http.ListenAndServe(":8080", nil)
}
