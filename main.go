package main

import (
	"log"
	"net/http"

	"github.com/TheGolurk/RestAPI/user"
)

func main() {
	http.HandleFunc("/user", user.HandleUserMethod)

	log.Println("Server started at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
