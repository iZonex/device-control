package main

import (
	"log"
	"net/http"
	"time"

	"github.com/iZonex/device-control/handlers"
	"github.com/iZonex/device-control/middleware"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static/"))
	r.HandleFunc("/", middleware.Chain(handlers.WifiHandler, middleware.Method("GET")))
	r.HandleFunc("/status", middleware.Chain(handlers.MainHandler, middleware.Method("GET")))
	r.HandleFunc("/api/status", middleware.Chain(handlers.DeviceInformationHandler, middleware.Method("GET")))
	r.HandleFunc("/wifi", middleware.Chain(handlers.WifiHandler, middleware.Method("GET")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	r.HandleFunc("/server", middleware.Chain(handlers.ServerHandler, middleware.Method("GET")))

	http.Handle("/", r)
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
