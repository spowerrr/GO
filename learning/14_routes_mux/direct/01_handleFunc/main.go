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
	http.HandleFunc("/dog", d)
	http.HandleFunc("/hog", h)
	http.ListenAndServe(":8080", nil)
}
