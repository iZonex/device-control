package main

import (
	"log"
	"net/http"
	"time"

	"github.com/iZonex/device-control/handlers"
	"github.com/iZonex/device-control/middleware"

	"github.com/gorilla/mux"
	"github.com/iZonex/device-control/handlers"
)

func main() {

	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static/"))
	r.HandleFunc("/", handlers.WifiHandler).Methods("GET")
	r.HandleFunc("/status", handlers.MainHandler).Methods("GET")
	r.HandleFunc("/api/status", handlers.DeviceInformationHandler).Methods("GET")
	r.HandleFunc("/wifi", handlers.WifiHandler).Methods("GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	r.HandleFunc("/server", handlers.ServerHandler).Methods("GET")

	http.Handle("/", r)
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
