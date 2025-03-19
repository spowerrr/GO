package main

import (
	"fmt"
	"log"
	"myapp/config"
	"myapp/routes"
	"net/http"
)

func main() {
	//connect to the database
	config.ConnectDB()

	//setup routes
	r := routes.SetupRoutes()

	//start the server
	fmt.Println("🚀 Server is running on port 8080")
	log.Fatalln(http.ListenAndServe(":8080", r))
}
