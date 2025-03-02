// reading JSON request data
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `{
    "username":"alice",
    "email":"alice@example.com"
}`
	// json.Unmarshal converts JSON into a Go data structure.
	// The map[string]string stores JSON data as key-value pairs.
	// This is commonly used in backend APIs to read request bodies.
	var user map[string]string
	err := json.Unmarshal([]byte(jsonData), &user) //-->transforms into go datastructure
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Username:", user["username"])
	fmt.Println("Email:", user["email"])
}
