package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
)

type Error struct {
	code    int32
	message string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users/user", getUser).Methods("GET")
	r.HandleFunc("/users/edit", editUser).Methods("POST")
	http.Handle("/", r)

	fmt.Println("Start listening on 8000")
	http.ListenAndServe(":8000", nil)
}

func getUser(writer http.ResponseWriter, request *http.Request) {
	userId := request.Header.Get("X-UserId")
	if userId == "" {
		error := Error {
			code:    0,
			message: "Not authorized",
		}
		writer.WriteHeader(401)
		json.NewEncoder(writer).Encode(error)
		return
	}

	getUserInfo(userId)
}

func getUserInfo(userId string) {

}

func editUser(writer http.ResponseWriter, request *http.Request) {
	userId := request.Header.Get("X-UserId")
	if userId == "" {
		error := Error {
			code:    0,
			message: "Not authorized",
		}
		writer.WriteHeader(401)
		json.NewEncoder(writer).Encode(error)
		return
	}
	//
}