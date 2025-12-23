package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type UserRequest struct {
	Msg string `json:"msg"`
}

func main() {
	http.HandleFunc("GET /time", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Time is %s", time.Now())
	})

	http.HandleFunc("POST /echo", func(w http.ResponseWriter, r *http.Request) {
		var ur UserRequest
		err := json.NewDecoder(r.Body).Decode(&ur)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(ur)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	})

	log.Println("HTTP server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
