package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Define the structs that match the JSON structure
type Address struct {
	Street string `json:street`
	City   string `json:city`
	State  string `json:state`
}

type Person struct {
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Age     int      `json:"age"`
	Address Address  `json:"address"`
	Hobbies []string `json:"hobbies"`
}

func main() {
	// Open the JSON file
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	// Read all bytes from the file into a slice of bytes
	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error Reading file:", err)
		return
	}

	var person Person
	// Unmarshal JSON data into the Person struct
	err = json.Unmarshal(data, &person)
	if err != nil {
		fmt.Println("Error Unmarshal file:", err)
		return
	}
	// Access and print the fields from the struct
	fmt.Println("Name:", person.Name)
	fmt.Println("Email:", person.Email)
	fmt.Println("Age:", person.Age)
	fmt.Println("Street:", person.Address.Street)
	fmt.Println("City:", person.Address.City)
	fmt.Println("State:", person.Address.State)
	fmt.Println("Hobbies:", person.Hobbies)
}
