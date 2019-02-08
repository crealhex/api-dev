package main

import (
	"log"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		us := getAll()

		w.Header().Set("Content-type", "text/plain")
		w.WriteHeader(http.StatusOK)
		log.Println(w.Write([]byte(us.String())))
	}
}