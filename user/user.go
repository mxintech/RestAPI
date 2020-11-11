package user

import (
	"encoding/json"
	"net/http"

	"github.com/TheGolurk/RestAPI/models"
)

var storage Storage

func init() {
	storage = make(map[string]models.Meta)
}

// Storage ....
type Storage map[string]models.Meta

// GetBody ....
func GetBody(r *http.Request) (models.User, error) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// CreateUser creates a new user
func (s Storage) CreateUser(w http.ResponseWriter, r *http.Request) {
	user, err := GetBody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error al leer el body"))
		return
	}
	s[user.Name] = user.Meta
	message := models.Message{
		Code:    200,
		Message: "Creado con exito",
	}

	JSON, err := json.Marshal(message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error al convertir el mensaje"))
		return
	}
	w.WriteHeader(message.Code)
	w.Write(JSON)
}

// DeleteUser deletes a user
func (s Storage) DeleteUser(w http.ResponseWriter, r *http.Request) {

}

// UpdateUser updates a user
func (s Storage) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

// GetUser get one user by name or all users
func (s Storage) GetUser(w http.ResponseWriter, r *http.Request) {

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
