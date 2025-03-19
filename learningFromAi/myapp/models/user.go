package models

// Define the User Model
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `josn:"email"`
}
