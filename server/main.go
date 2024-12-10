package main

import (
	"fmt"
	"go-app/db"
	"go-app/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	db.CreateTable()

	router := mux.NewRouter()
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
