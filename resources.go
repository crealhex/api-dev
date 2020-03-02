package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	"strconv"
)

func friendsManager(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	if user == "" {
		http.Error(w, "missing user name in query string", http.StatusBadRequest)
		return
	}
	add := r.URL.Query().Get("add")
	if add == "" {
		http.Error(w, "missing add name in query string", http.StatusBadRequest)
		return
	}

	uno, _ := strconv.Atoi(user)
	dos, _ := strconv.Atoi(add)

	addFriend(uno, dos)
}

func solo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Error(w, "Unable to access", http.StatusBadRequest)
	}
	userID, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
	if r.Method == http.MethodPut {
		decoder := json.NewDecoder(r.Body)
		var uu User
		err := decoder.Decode(&uu)
		if err != nil {
			http.Error(w, "Error actualizando", http.StatusBadRequest)
			return
		}
		user := &User{
			userID,
			uu.Username,
			uu.Fullname,
			uu.Age,
			uu.Email,
			uu.City,
			uu.Phone,
			nil,
		}
		update(userID, user)
	}
	if r.Method == http.MethodDelete {
		deleteu(userID)
	}
	//user, _ := json.Marshal(getByID(userID))
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	//log.Println(w.Write(user))
}

// nu defines: new user
func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "No encontrado", http.StatusNotFound)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var nu User
	err := decoder.Decode(&nu)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	user := &User{
		0,
		nu.Username,
		nu.Fullname,
		nu.Age,
		nu.Email,
		nu.City,
		nu.Phone,
		nil,
	}
	add(user)
	ok := map[string]string{
		"response": "El usuario ha sido creado correctamente",
		"code": "200",
	}
	w.Header().Set("Content-Type", "application/json")
	log.Println(json.NewEncoder(w).Encode(ok))
}

func users(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "No encontrado", http.StatusNotFound)
		return
	}
	us, _ := json.Marshal(getAll())
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println(w.Write(us))
}