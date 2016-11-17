/*
Package db is a Go library that populates sql databases with Bro log data.
It was built to work in conjunction with the github.com/amadeovezz/gobro/parse library.
*/
package db

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// db represents a valid connection the the database.
var db *sql.DB

// GetDBConn returns a db connection.
func GetDBConn() *sql.DB {
	return db
}

// InitDB ensures that a valid connection to the database is established.
func InitDB(user, pw, ip, port, dbase string) error {

	sqlConn, err := ConnectToSql(user, pw, ip, port, dbase)
	if err != nil {
		return err
	}

	db = sqlConn

	return nil
}

// ConnectToSql tries to connect to the database for a total of 28 seconds.
// Everytime it cannot connect to the db, it sleeps for +1 seconds longer than the
// previous iteration. Any other errors besides "connection refused errors" are returned.
func ConnectToSql(user string, pw string, ip string, port string, dbase string) (*sql.DB, error) {

	// This call doesn't actually communicate with the db
	// it just checks if arguments are valid
	db, err := sql.Open("mysql", user+":"+pw+"@("+ip+":"+port+")/"+dbase)
	if err != nil {
		return nil, err
	}

	var currentTime int = 1
	const maxTime int = 7

	for {
		err = db.Ping()

		if currentTime == maxTime {
			return nil, errors.New("Connection to mysql timed out")
		} else if err != nil {
			if strings.Contains(err.Error(), "connection refused") {
				fmt.Println("Couldnt connect to mysql, retrying in ", currentTime, " seconds")
				time.Sleep(time.Duration(currentTime) * time.Second)
				currentTime++
				continue
			} else {
				return nil, err
			}

		} else {
			break
		}
	}

	return db, nil

}

// InsertBatch, reads from a channel of values and inserts them into the db.
func InsertBatch(values chan []string, logType string, numOfValues int) error {

	tx, err := db.Begin()

	if err != nil {
		return err
	}

	insert := "INSERT INTO " + logType + " VALUES (?" + strings.Repeat(",?", numOfValues-1) + ")"

	stmt, err := db.Prepare(insert)

	if err != nil {
		tx.Rollback()
		return err
	}

	for record := range values {

		// convert slice to contain interface types
		newRecord := make([]interface{}, len(record))
		for i, v := range record {
			newRecord[i] = v
		}

		_, err = stmt.Exec(newRecord...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil

}
