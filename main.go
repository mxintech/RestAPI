package main

import (
	"log"
	"net/http"

	"github.com/TheGolurk/RestAPI/user"
)

func main() {

	http.HandleFunc("/user", user.HandleUser)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
