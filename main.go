package main

import (
	"apirest/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Rutas
	mux := mux.NewRouter()

	// Endpoind
	mux.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/user/{id}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+} ", handlers.DeleteUser).Methods("DELETE")

	fmt.Println("Hola")
	log.Println("Servidor Corriendo")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
