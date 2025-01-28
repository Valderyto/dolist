package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Done     bool   `json:"done"`
	Assignee *User  `json:"assignee,omitempty"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var (
	tasks  []Task
	mutex  = &sync.Mutex{}
	taskID = 1
)

func getTasks(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	task.ID = taskID
	taskID++
	tasks = append(tasks, task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func main() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getTasks(w, r)
			return
		}
		if r.Method == http.MethodPost {
			createTask(w, r)
			return
		}
		http.NotFound(w, r)
	})

	http.ListenAndServe(":8001", nil)
}
