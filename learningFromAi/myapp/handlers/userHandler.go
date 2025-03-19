// Create Handlers for users
package handlers

import (
	"encoding/json"
	"myapp/config"
	"myapp/models"
	"net/http"
)

// Create a new user
func CreateUser(w http.ResponseWriter, req *http.Request) {
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	query := `INSERT INTO users (name,email) VALUES ($1,$2) RETURNING id`
	err = config.DB.QueryRow(query, user.Name, user.Email).Scan(&user)
	if err != nil {
		http.Error(w, "Error inserting user:", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
//get all users
func GetUsers(w http.ResponseWriter, req *http.Request){
	rows,err:=config.DB.Query("SELECT id,name,email FROM users")
	if err!=nil{
		http.Error(w,"Error retriving users",http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next(){
		var user models.User
		err:=rows.Scan(&user.ID,&user.Name,&user.Email)
		if err!=nil{
			http.Error(w,"Error scanning users",http.StatusInternalServerError)
			return
		}
		users =append(users,user)
	}
	json.NewEncoder(w).Encode(users)
}