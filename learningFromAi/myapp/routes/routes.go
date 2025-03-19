package routes

import (
	"myapp/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router{
	r:=mux.NewRouter()
	r.HandleFunc("/users",handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users",handlers.GetUsers).Methods("GET")
	return r
}