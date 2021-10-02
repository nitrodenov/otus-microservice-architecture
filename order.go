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
	ID      string `json:"id"`
	OwnerId string `json:"ownerId"`
	Price   int64  `json:"price"`
	Version int64  `json:"version"`
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
		fmt.Println("Error in order decode")
	}
	orderId := uuid.New().String()
	order := Order{
		ID:      orderId,
		OwnerId: userId,
		Price:   orderRequest.Price,
		Version: version,
	}
	latestVersion, err := createOrder(order)
	if err != nil {
		fmt.Println("Error in order createOrder")
		writer.Header().Add("ETag", strconv.FormatInt(latestVersion, 10))
		writer.WriteHeader(409)
		return
	}

	writer.Header().Add("ETag", strconv.FormatInt(latestVersion, 10))
	json.NewEncoder(writer).Encode(order)
	writer.WriteHeader(200)
}

func orders(writer http.ResponseWriter, request *http.Request) {
	userId := request.Header.Get("X-UserId")
	orders := getOrdersByOwnerId(userId)

	version := strconv.FormatInt(orders[0].Version, 10)
	writer.Header().Add("ETag", version)
	json.NewEncoder(writer).Encode(orders)
	writer.WriteHeader(200)
}

func createOrder(order Order) (int64, error) {
	db := createConnection()
	defer db.Close()

	latestVersion := getLatestVersionOfOrderListForOwner(order.OwnerId)
	if latestVersion != order.Version {
		return latestVersion, nil
	}

	var id string
	sqlStatement := `INSERT INTO orders (id, owner_id, price, version) VALUES ($1, $2, $3, $4) ON CONFLICT (owner_id, version) DO NOTHING RETURNING Id`
	err := db.QueryRow(sqlStatement, order.ID, order.OwnerId, order.Price, order.Version+1).Scan(&id)
	if err != nil {
		version := getLatestVersionOfOrderListForOwner(order.OwnerId)
		fmt.Println("Error in createOrder")
		return version, err
	}
	return order.Version, nil
}

//func updateOrderStatusById(orderId string, orderStatus string) {
//	db := createConnection()
//	defer db.Close()
//
//	var id string
//	sqlStatement := `UPDATE orders SET status = $1 WHERE id = $2`
//	err := db.QueryRow(sqlStatement, orderStatus, orderId).Scan(&id)
//	if err != nil {
//		fmt.Println("Error in updateOrderStatusById")
//	}
//}

func getLatestVersionOfOrderListForOwner(ownerId string) int64 {
	db := createConnection()
	defer db.Close()

	var version string
	sqlStatement := `SELECT version AS latest_version from orders where owner_id = $1 order by version desc limit 1`
	err := db.QueryRow(sqlStatement, ownerId).Scan(&version)
	if err != nil {
		fmt.Println("Error in getLatestVersionOfOrderListForOwner")
		return -1
	}

	vers, err := strconv.ParseInt(version, 10, 64)
	if err != nil {
		return 0
	}
	return vers
}

func getOrdersByOwnerId(userId string) []Order {
	db := createConnection()
	defer db.Close()

	var orders []Order
	sqlStatement := `SELECT id, owner_id, price, version FROM orders WHERE owner_id = $1 ORDER BY version DESC`
	err := db.QueryRow(sqlStatement, userId).Scan(&orders)
	if err != nil {
		fmt.Println("Error in getOrdersByOwnerId")
		return nil
	}

	return orders
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
