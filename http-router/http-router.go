package httprouter

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matellio-golang-test/config"
	"github.com/matellio-golang-test/handler"
)

func StartHTTPServer() (err error) {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.GetAllPorts).Methods("GET")
	r.HandleFunc("/api/ports", handler.GetAllPorts).Methods("GET")
	r.HandleFunc("/api/port/{id}", handler.GetPortByID).Methods("GET")
	r.HandleFunc("/api/port/{id}", handler.CreateNewPort).Methods("POST")
	r.HandleFunc("/api/port/{id}", handler.UpdatePort).Methods("PUT")
	r.HandleFunc("/api/port/{id}", handler.DeletePortByID).Methods("DELETE")
	http.Handle("/", r)
	httpErr := http.ListenAndServe(config.GetHTTPServerAddress(), r)
	if httpErr != nil {
		log.Fatalln("http error ", httpErr)
	}
	return
}
