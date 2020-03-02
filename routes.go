package main

import (
	"fmt"
	"log"
	"net/http"
)

func startRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/register", register)
	http.HandleFunc("/users", users)
	http.HandleFunc("/users/", solo)
	http.HandleFunc("/friends", friendsManager)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Fprintf(w, "<h1>NO NEW FRIENDS API</h1>"))
}