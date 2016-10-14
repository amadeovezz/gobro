package db

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB(user, pw, ip, port, dbase string) error {

	mySql, err := ConnectMySql(user, pw, ip, port, dbase)
	if err != nil {
		return err
	}

	db = mySql

	return nil
}

func ConnectMySql(user string, pw string, ip string, port string, dbase string) (*sql.DB, error) {

	// This call doesn't actually communicate with the db
	// it just checks if arguments are valid
	db, err := sql.Open("mysql", user+":"+pw+"@("+ip+":"+port+")/"+dbase)
	if err != nil {
		return nil, err
	}

	// This will wait for a total of 28 seconds before it gives up
	// 1 + 2 + 3 + 4 + 5 + 6 + 7
	var second int = 1
	const maxTime int = 7

	for {
		err = db.Ping()

		if err != nil {
			if strings.Contains(err.Error(), "connection refused") {
				fmt.Println("Couldnt connect to mysql, retrying in ", second, " seconds")
				time.Sleep(time.Duration(second) * time.Second)
				second++
				continue
			} else {
				return nil, err
			}

		} else if second == maxTime {
			return nil, errors.New("Connection to mysql timed out")
		} else {
			break
		}
	}

	return db, nil

}

type Conn struct {
	time       uint64
	connUID    string
	origIp     string
	origPort   string
	respIp     string
	respPort   string
	proto      string
	service    string
	duration   float64
	inBytes    uint64
	outBytes   uint64
	inPackets  uint64
	outPackets uint64
}

func InsertBatchIntoConn(connRecord []Conn) error {

	tx, err := db.Begin()

	if err != nil {
		return err
	}

	stmt, err := db.Prepare(`
		INSERT INTO conn
		(time, conn_uid, orig_ip, orig_port, resp_ip, resp_port, proto,
		service, duration, in_bytes, out_bytes, in_packets, out_packets)
		VALUES
		(?,?,?,?,?,?,?,?,?,?,?,?,?)
	`)

	if err != nil {
		tx.Rollback()
		return err
	}

	for _, conn := range connRecord {
		_, err := stmt.Exec(
			conn.time,
			conn.connUID,
			conn.origIp,
			conn.origPort,
			conn.respIp,
			conn.respPort,
			conn.proto,
			conn.service,
			conn.duration,
			conn.inBytes,
			conn.outBytes,
			conn.inPackets,
			conn.outPackets,
		)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil

}
