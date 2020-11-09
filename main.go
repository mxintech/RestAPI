package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/user", User)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func User(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Get users"))
	}
}
