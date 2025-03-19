// Middleware functions allow you to execute logic before or after handling a request. Common use cases:
//
// Logging requests
// Authentication/Authorization
// Request validation

package main

import (
	"fmt"
	"log"
	"net/http"
)

//Maddleware function seperately used in the 2 comments function
// Logging function that acts as middleware
// func logRequest(w http.ResponseWriter, req *http.Request, next http.Handler) {
// 	log.Printf("Request: %s %s", req.Method, req.URL.Path)
// 	next.ServeHTTP(w, req) // Call the next handler
// }
//
// // Middleware function that wraps a handler with logging
// func loggingMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
// 		logRequest(w, req, next)
// 	})
// }

// Middleware function to log requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Request: %s %s", req.Method, req.URL.Path)
		next.ServeHTTP(w, req) //call the next handler in the chain
	})
}
func hellowHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "HI,Backend")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/middle", hellowHandler)
	//apply middleware
	loggedMux := loggingMiddleware(mux)

	log.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", loggedMux)
	if err != nil {
		log.Fatalln("Error reaching in port 8080:", err)
	}
}
