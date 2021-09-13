package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

type User struct {
	id        string `json:"id"`
	login     string `json:"login"`
	password  string `json:"password"`
	email     string `json:"email"`
	firstName string `json:"firstName"`
	lastName  string `json:"lastName"`
}

var sessions map[string]User

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", register).Methods("POST")
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/signin", signin).Methods("GET")
	r.HandleFunc("/auth", auth).Methods("GET")
	r.HandleFunc("/logout", logout).Methods("GET")
	http.Handle("/", r)

	fmt.Println("Start listening on 8000")
	http.ListenAndServe(":8000", nil)
}

func register(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	writer.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "POST")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user User

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Error in register")
	}

	insertUser(user)
	writer.WriteHeader(200)
}

func login(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	writer.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "POST")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user User

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Error in login")
	}

	userInfo, err := getUserInfo(user.login, user.password)
	if err != nil {
		log.Fatalf("Error in login after getting user info")
	}

	sessionId := createSession(user)
	http.SetCookie(writer, &http.Cookie{
		Name:     "session_id",
		Value:    sessionId,
		HttpOnly: true,
	})
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(userInfo)
}

func signin(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode("{message: Please go to login and provide Login/Password}")
}

func auth(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("session_id")
	if err != nil {
		writer.WriteHeader(401)
		return
	}
	user := sessions[cookie.Value]
	writer.Header().Add("X-UserId", user.id)
	writer.Header().Add("X-User", user.login)
	writer.Header().Add("X-Email", user.email)
	writer.Header().Add("X-First-Name", user.firstName)
	writer.Header().Add("X-Last-Name", user.lastName)
}

func logout(writer http.ResponseWriter, request *http.Request) {
	http.SetCookie(writer, &http.Cookie{
		Name:    "session_id",
		Value:   "",
		Expires: time.Time{},
	})
}

func getUserInfo(login string, password string) (User, error) {
	db := createConnection()
	defer db.Close()

	var user User

	sqlStatement := `SELECT * FROM users WHERE login=$1 AND password=$2`
	row := db.QueryRow(sqlStatement, login, password)
	err := row.Scan(&user.login, &user.password, &user.email, &user.firstName, &user.lastName)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return user, err
}

func insertUser(user User) int64 {
	db := createConnection()
	defer db.Close()

	userId := uuid.New().String()
	sqlStatement := `INSERT INTO users (id, login, password, email, firstName, lastName) VALUES ($1, $2, $3, $4, $5, $6) RETURNING Id`

	var id int64

	err := db.QueryRow(sqlStatement, userId, user.login, user.password, user.email, user.firstName, user.lastName).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)
	return id
}

func createSession(user User) string {
	sessionId := uuid.New().String()
	sessions[sessionId] = user
	return sessionId
}

func createConnection() *sql.DB {
	psqlconn := os.Getenv("DATABASE_URI")
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}
