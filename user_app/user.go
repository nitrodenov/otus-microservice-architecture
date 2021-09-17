package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

type User struct {
	Id        string `json:"id"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users/user", getUser).Methods("GET")
	r.HandleFunc("/users/edit", edit).Methods("POST")
	http.Handle("/", r)

	fmt.Println("Start listening on 8000")
	http.ListenAndServe(":8000", nil)
}

func getUser(writer http.ResponseWriter, request *http.Request) {
	userId := request.Header.Get("X-UserId")
	if userId == "" {
		error := Error{
			Code:    0,
			Message: "Not authorized. getUser",
		}
		writer.WriteHeader(401)
		json.NewEncoder(writer).Encode(error)
		return
	}

	userInfo, err := getUserInfo(userId)
	if err != nil {
		log.Fatalf("Error in get user")
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(userInfo)
}

func edit(writer http.ResponseWriter, request *http.Request) {
	userId := request.Header.Get("X-UserId")
	if userId == "" {
		error := Error{
			Code:    0,
			Message: "Not authorized",
		}
		writer.WriteHeader(401)
		json.NewEncoder(writer).Encode(error)
		return
	}

	var user User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Error in edit user")
	}

	if user.Id != userId {
		log.Fatalf("Error in edit user. user.id != userId")
	}

	editUser(user)
	writer.WriteHeader(200)
}

func editUser(user User) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := `UPDATE users SET login=$2, firstName=$3, email=$4, firstName=$5, lastName=$6 WHERE id=$1`

	res, err := db.Exec(sqlStatement, user.Id, user.Login, user.Email, user.FirstName, user.LastName)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected
}

func getUserInfo(userId string) (User, error) {
	db := createConnection()
	defer db.Close()

	var user User

	sqlStatement := `SELECT * FROM users WHERE id=$1`
	row := db.QueryRow(sqlStatement, userId)
	err := row.Scan(&user.Login, &user.Password, &user.Email, &user.FirstName, &user.LastName)

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
