package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Notification struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	UserID  int    `json:"user_id"`
}

var (
	notifications  []Notification
	notificationID = 1
	mutex          = &sync.Mutex{}
)

func getNotifications(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifications)
}

func createNotification(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	var notification Notification
	json.NewDecoder(r.Body).Decode(&notification)
	notification.ID = notificationID
	notificationID++
	notifications = append(notifications, notification)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(notification)
	fmt.Printf("Notification sent to user %d: %s\n", notification.UserID, notification.Message)
}

func main() {
	http.HandleFunc("/notifications", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getNotifications(w, r)
			return
		}
		if r.Method == http.MethodPost {
			createNotification(w, r)
			return
		}
		http.NotFound(w, r)
	})

	http.ListenAndServe(":8002", nil)
}
