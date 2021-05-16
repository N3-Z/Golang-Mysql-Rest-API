package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	connectStr string
}

const (
	dbusername string = "root"
	dbpassword string = ""
	dbdatabase string = "goTesting"
	dbaddress  string = "localhost"
	dbport     string = "3306"
)

var (
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbusername, dbpassword, dbaddress, dbport, dbdatabase)
)

func (db *Database) mysqlConnect() (*sql.DB, error) {
	db.connectStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbusername, dbpassword, dbaddress, dbport, dbdatabase)
	conn, err := sql.Open("mysql", db.connectStr)

	if err != nil {
		return nil, err
	}

	return conn, nil
}
