package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)

	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `,!-- image dosen't serve--> <img src="/toby.jpg"`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat() // Retrieves file info (size, modification time, etc.)
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	// Serve the file efficiently with its modification time
	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}

// Why Use http.ServeContent() Instead of io.Copy()?
// Supports HTTP caching (browsers can use the Last-Modified timestamp).
// Supports byte-range requests (for partial downloads, especially useful for videos or large images).
// Automatically sets headers (like Content-Type).

// Better Alternative: Use http.ServeFile()
// Instead of manually opening and serving the file, you can simplify:

// func dogPic(w http.ResponseWriter, req *http.Request) {
// 	http.ServeFile(w, req, "toby.jpg")
// }

// Automatically handles everything (headers, errors, caching).
// More efficient and cleaner.
