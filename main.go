package main

import (
	"encoding/json"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello world")
	})

	http.ListenAndServe(":8081", nil)
}
