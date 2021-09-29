package otus

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"strconv"
)

type Order struct {
	ID          string `json:"id"`
	OwnerId     string `json:"ownerId"`
	Price       int64  `json:"price"`
	OrderStatus string `json:"orderStatus"`
	Version     int64  `json:"version"`
}

type OrderRequest struct {
	UserId string `json:"userId"`
	Price  int64  `json:"price"`
}

type OrderResponse struct {
	ID        string `json:"id"`
	IsSuccess bool   `json:"isSuccess"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/order", order).Methods("POST")
	r.HandleFunc("/orders", orders).Methods("GET")
	http.Handle("/", r)

	fmt.Println("Start listening on 8000")
	http.ListenAndServe(":8000", nil)
}

func order(writer http.ResponseWriter, request *http.Request) {
	userId := request.Header.Get("X-UserId")
	version, _ := strconv.ParseInt(request.Header.Get("If-Match"), 10, 64)
	var orderRequest OrderRequest
	err := json.NewDecoder(request.Body).Decode(&orderRequest)
	if err != nil {

	}
	orderId := uuid.New().String()
	order := Order{
		ID:          orderId,
		OwnerId:     userId,
		Price:       orderRequest.Price,
		OrderStatus: "IN_PROGRESS",
		Version:     version,
	}
	latestVersion, err := createOrder(order)
	if err != nil {

	}

	updateOrderStatusById(orderId, "SUCCESS")

	writer.Header().Add("ETag", strconv.FormatInt(latestVersion, 10))
	json.NewEncoder(writer).Encode(order)
	writer.WriteHeader(200)
}

func orders(writer http.ResponseWriter, request *http.Request) {

}

func createOrder(order Order) (int64, error) {
	db := createConnection()
	defer db.Close()

	latestVersion := getLatestVersionOfOrderListForOwner(order.OwnerId)
	if latestVersion != order.Version {
		return latestVersion, nil
	}

	var id string
	sqlStatement := `INSERT INTO orders (id, owner_id, price, status, version) VALUES ($1, $2, $3, $4, $5) ON CONFLICT (owner_id, version) DO NOTHING RETURNING Id`
	err := db.QueryRow(sqlStatement, order.ID, order.OwnerId, order.Price, order.OrderStatus, order.Version+1).Scan(&id)
	if err != nil {
		version := getLatestVersionOfOrderListForOwner(order.OwnerId)
		return version, err
	}
	return order.Version, nil
}

func updateOrderStatusById(orderId int64, orderStatus string) {
	db := createConnection()
	defer db.Close()

	var id string
	sqlStatement := `UPDATE orders SET status = $1 WHERE id = $2`
	err := db.QueryRow(sqlStatement, orderStatus, orderId).Scan(&id)
	if err != nil {

	}
}

func getLatestVersionOfOrderListForOwner(ownerId string) int64 {
	db := createConnection()
	defer db.Close()

	var version string
	sqlStatement := `SELECT version AS latest_version from orders where owner_id = $1 order by version desc limit 1`
	err := db.QueryRow(sqlStatement, ownerId).Scan(&version)
	if err != nil {

	}

	vers, err := strconv.ParseInt(version, 10, 64)
	if err != nil {
		return 0
	}
	return vers
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
