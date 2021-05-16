package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	db        = new(Database)
	conn, err = db.mysqlConnect()
)

func handleHome(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}

func handleRequest() {
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[*] Success")

	router := mux.NewRouter()
	router.HandleFunc("/", handleHome).Methods("GET")
	router.HandleFunc("/user/", getAllUser).Methods("GET")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))

	conn.Close()
	fmt.Println("[*] Exit")
}

func main() {
	fmt.Println("Server Start")
	handleRequest()
}
