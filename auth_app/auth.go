package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Session struct {
	id        int
	login     string
	email     string
	firstName string
	lastName  string
}

type Cred struct {
	login string
	password string
}

var sessions map[string]Session

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/signin", signin).Methods("GET")
	r.HandleFunc("/register", register).Methods("GET")
	r.HandleFunc("/login", login).Methods("GET")
	r.HandleFunc("/auth", auth).Methods("GET")
	r.HandleFunc("/logout", logout).Methods("GET")
	http.Handle("/", r)

	fmt.Println("Start listening on 8000")
	http.ListenAndServe(":8000", nil)
}

func signin(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode("{message: Please go to login and provide Login/Password}")
}

func register(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {

	}
	var session Session
	err = json.Unmarshal(body, session)
	if err != nil {

	}


}

func login(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {

	}
	var cred Cred
	err = json.Unmarshal(body, cred)
	if err != nil {

	}

	userInfo := getUserInfo(cred.login, cred.password)
	if userInfo == nil {

	}
}

func getUserInfo(l string, password string) interface{} {

}

func auth(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("session_id")
	if err != nil {
		writer.WriteHeader(401)
		return
	}
	session := sessions[cookie.Value]
	writer.Header().Add("X-UserId", strconv.Itoa(session.id))
	writer.Header().Add("X-User", session.login)
	writer.Header().Add("X-Email", session.email)
	writer.Header().Add("X-First-Name", session.firstName)
	writer.Header().Add("X-Last-Name", session.lastName)
}

func logout(writer http.ResponseWriter, request *http.Request) {
	http.SetCookie(writer, &http.Cookie{
		Name:    "session_id",
		Value:   "",
		Expires: time.Time{},
	})
}
