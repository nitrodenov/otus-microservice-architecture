package order

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/orders", orders).Methods("POST")
	r.HandleFunc("/orders/declines/{orderId}", declines).Methods("PUT")
	r.HandleFunc("/orders/approvals", approvals).Methods("PUT")
	http.Handle("/", r)

	fmt.Println("Start listening on 8000")
	http.ListenAndServe(":8000", nil)
}

func orders(writer http.ResponseWriter, request *http.Request) {

}

func declines(writer http.ResponseWriter, request *http.Request) {

}

func approvals(writer http.ResponseWriter, request *http.Request) {

}
