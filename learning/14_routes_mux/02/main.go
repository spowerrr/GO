package main

import (
	"io"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "This is dog page")
	case "/hot":
		io.WriteString(w, "This is hot page")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
