package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/health", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`{"status":"OK"}`))
}
