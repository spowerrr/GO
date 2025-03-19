package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/bog", bog)
	http.ListenAndServe(":8080", nil)
}

func bog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8") // This line sets the response header to indicate that the response body is HTML (instead of plain text). The charset=utf-8 ensures the text encoding is UTF-8, which supports special characters.
	io.WriteString(w, `<!-- not serving from our server -->
  <img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">`)
}
