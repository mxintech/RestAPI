package user

import "net/http"

// HandleUser handle method for request /user
func HandleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Want's to request all users? :)"))
	}
}
