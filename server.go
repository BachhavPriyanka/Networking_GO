package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleHome)

	fmt.Println("Listening on :8080...")
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
