package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/info", handleInfo).Methods("GET")

	r.HandleFunc("/test/{item}", handleTest).Methods("GET").Queries("sc", "{sc}")
	r.HandleFunc("/test/{item}", handleTest).Methods("GET")

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	srv := &http.Server{
		Handler:      loggedRouter,
		Addr:         ":8088",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Started Server")

	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf(err.Error())
	}
}

func handleInfo(w http.ResponseWriter, _ *http.Request) {
	reply, _ := json.Marshal("it works V3!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(reply)
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	var body []byte
	r.Body.Read(body)

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}

	reply, _ := json.MarshalIndent(
		struct {
			Body json.RawMessage
			Vars map[string]string
			Url  string
			Dump string
		}{
			Vars: mux.Vars(r),
			Body: body,
			Url:  r.URL.String(),
			Dump: string(requestDump),
		}, "", "    ")

	fmt.Println(string(requestDump))
	fmt.Printf("Reply: %s\n", string(reply))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(reply)
}
