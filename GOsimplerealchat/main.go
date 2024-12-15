package main

import (
	"fmt"
	"gosimplerealchat/controller"
	"gosimplerealchat/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting the server...")

	var r = mux.NewRouter()

	routes.RegisterRoutes(r)
	// Serve static files from /static/ URL path
	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/").Handler(fs)
	// Register your routes
	go controller.HandleMessages()

	// Start the server
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
