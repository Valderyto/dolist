package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var (
	users  []User
	mutex  = &sync.Mutex{}
	userID = 1
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	user.ID = userID
	userID++
	users = append(users, user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getUsers(w, r)
			return
		}
		if r.Method == http.MethodPost {
			createUser(w, r)
			return
		}
		http.NotFound(w, r)
	})

	http.ListenAndServe(":8000", nil)
}
