package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	// menggunakan prepare statement agar dapat mengatasi serangan injection
	stmt, err := conn.Prepare("select id, email, username from users where id = ?")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer stmt.Close()

	var x = User{}

	stmt.QueryRow(id).Scan(&x.ID, &x.Email, &x.Username)

	byte_data, _ := json.Marshal(x)
	writer.Write(byte_data)
}

func getAllUser(writer http.ResponseWriter, request *http.Request) {
	stmt, err := conn.Query("select id, email, username from users")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer stmt.Close()

	var users []User

	for stmt.Next() {
		var x User
		stmt.Scan(&x.ID, &x.Email, &x.Username)
		users = append(users, x)
	}

	byte_data, _ := json.Marshal(users)

	writer.Write(byte_data)
}
