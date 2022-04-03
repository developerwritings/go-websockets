package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	log.Println("Connected to Server")
	http.ListenAndServe(":8080", mux)
}
