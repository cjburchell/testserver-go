package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/info", handleInfo).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8088",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf(err.Error())
	}
}

func handleInfo(w http.ResponseWriter, _ *http.Request) {
	reply, _ := json.Marshal("it works!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(reply)
}
