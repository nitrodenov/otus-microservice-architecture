package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/items/reservations", deliveries).Methods("POST")
	r.HandleFunc("/items/reservations/cancellations/{reservationId}", cancellations).Methods("POST")
	http.Handle("/", r)

	fmt.Println("Start listening on 8000")
	http.ListenAndServe(":8000", nil)
}

func deliveries(writer http.ResponseWriter, request *http.Request) {

}

func cancellations(writer http.ResponseWriter, request *http.Request) {

}
