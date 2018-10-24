package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func handleInfo(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/info", handleInfo).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8082",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf(err.Error())
	}
}
