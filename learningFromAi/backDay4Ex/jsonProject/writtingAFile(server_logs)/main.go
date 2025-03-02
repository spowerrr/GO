package main

import (
	"fmt"
	"os"
)

func main() {
	log := "New Request: GET/users" // this is just a message witten in the server.log file
	os.WriteFile("server.log", []byte(log), 0644)
	fmt.Println("Log saved successfully")
}
