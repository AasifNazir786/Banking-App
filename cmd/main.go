package main

import (
	"Go-GitHub-Projects/Banking-App/handlers"
	"net/http"
)

func main() {

	http.HandleFunc("/create-account", handlers.CreateAccount)
}
