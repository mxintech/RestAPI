package user

import (
	"net/http"
)

var storage Storage

func init() {
	storage = make(map[string]*User)
}

type Storage map[string]*User

// User...
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// CreateUser creates a new user
func (U *Storage) CreateUser(w http.ResponseWriter, r *http.Request) {

}

// DeleteUser deletes a user
func (U *Storage) DeleteUser(w http.ResponseWriter, r *http.Request) {

}

// UpdateUser updates a user
func (U *Storage) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

// GetUser get one user by name or all users
func (U Storage) GetUser(w http.ResponseWriter, r *http.Request) {

}

// HandleUser handle method for request /user
func HandleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		storage.GetUser(w, r)
		return
	case "POST":
		storage.CreateUser(w, r)
		return
	case "PUT":
		storage.UpdateUser(w, r)
		return
	case "DELETE":
		storage.DeleteUser(w, r)
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
