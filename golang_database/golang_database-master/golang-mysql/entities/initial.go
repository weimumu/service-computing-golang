package entities

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var mydb *sql.DB

func init() {
	//https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	db, err := sql.Open("mysql", "test:test2017@tcp(139.199.174.146:3306)/mytest?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	} else {
		fmt.Print("successfully connect to DataBase")
	}
	mydb = db
}

// SQLExecer interface for supporting sql.DB and sql.Tx to do sql statement
type SQLExecer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// DaoSource Data Access Object Source
type DaoSource struct {
	// if DB, each statement execute sql with random conn.
	// if Tx, all statements use the same conn as the Tx's connection
	SQLExecer
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
