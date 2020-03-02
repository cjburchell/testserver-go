package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func handleInfo(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write([]byte("This is a test"))
	if err != nil{
		fmt.Println(err.Error())
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/info", handleInfo).Methods("GET")

	port:=8082
	fmt.Printf("Starting Server at port %d\n", port)
	srv := &http.Server{
		Handler:      r,
		Addr:          fmt.Sprintf(":%d", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println(err.Error())
		}
		fmt.Println("http server shut down")
	}()

	defer stopHTTPServer(srv)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Println("Shutting down")
	os.Exit(0)
}

func stopHTTPServer(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	fmt.Println("Shutting down http server")
	err := srv.Shutdown(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}
}
