package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UsersDataWrapper struct {
	Info User `json:"users,omitempty"`
}

type User struct {
	ID int32 `json:"user_id,omitempty"`
	Nombre string `json:"nombre,omitempty"`
	Apellido string `json:"apellido,omitempty"`
	Edad uint8 `json:"edad,omitempty"`
	Friends []User `json:"friends,omitempty"`
}

func getJsonResponse()  ([]byte, error) {
	friend := User{
		13,
		"Luis",
		"Enco",
		20,
		nil,
	}

	friends := []User{
		friend,
	}

	user := User{
		1,
		"Dayan",
		"No sé",
		19,
		friends,
	}

	system := UsersDataWrapper{user}

	return json.MarshalIndent(system, "", "  ")
}

//func serveRest(w http.ResponseWriter, r *http.Request) {
//	response, err := getJsonResponse()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Fprintf(w, string(response))
//}

func customHFM(message string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, message)
		},
	)
}

var userstore = make(map[string]User)

func setUser(w http.ResponseWriter, r *http.Request) http.Handler {
	var users []User
	for _, v := range userstore {
		users = append(users, v)
	}
	w.Header().Set("Content-Type", "application/json")

	j, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}


func main() {
	mux := http.NewServeMux()

	response, err := getJsonResponse()
	if err != nil {
		panic(err)
	}

	mux.Handle("/users", customHFM(string(response)))
	mux.Handle("/users", customHFM(string(response)))
	//mux.Handle("/users/", getUsers)

	log.Println("Ejecutando...")
	log.Println(http.ListenAndServe(":8080", mux))
}