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

func (U *Storage) CreateUser(w http.ResponseWriter, r *http.Request) {

}

func (U *Storage) DeleteUser(w http.ResponseWriter, r *http.Request) {

}

func (U *Storage) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (U Storage) GetUser(w http.ResponseWriter, r *http.Request) {

}

// HandleUser handle method for request /user
func HandleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Want's to request all users? :)"))
	case "POST":

	case "PUT":

	case "DELETE":

	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
