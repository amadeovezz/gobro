package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/amadeovezz/gobro/config"
)

var conf config.Config

func TestMain(m *testing.M) {
	conf.SetupConfig("config.toml")

	err := InitDB(
		conf.DB.Username,
		conf.DB.Password,
		conf.DB.IP,
		conf.DB.Port,
		conf.DB.DatabaseName,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	m.Run()

}

func TestBatchInsert(t *testing.T) {

	// Create buffer
	connBuffer := make(chan []string, 20)

	record := []string{"1476454019", "axb3912eK345", "10.0.0.20", "22", "8.8.8.8",
		"88", "tcp", "facebook", "10.001", "200", "4000", "30", "400"}

	connBuffer <- record

	record = []string{"1476454019", "bxb3912eK345", "10.0.0.20", "22", "8.8.8.8",
		"88", "tcp", "facebook", "10.001", "200", "4000", "30", "400"}

	connBuffer <- record

	close(connBuffer)

	// Insert into db
	err := InsertBatch(connBuffer, "conn", len(record))
	if err != nil {
		t.Error(err)
	}

	// Scan row
	broId := "axb3912eK345"
	var dbId string
	err = db.QueryRow("SELECT * FROM conn WHERE uid = ?", broId).Scan(dbId)

	if err == sql.ErrNoRows {
		t.Error("Row was not inserted properly")
	}

	// Scan row
	broId = "bxb3912eK345"
	err = db.QueryRow("SELECT * FROM conn WHERE uid = ?", broId).Scan(dbId)

	if err == sql.ErrNoRows {
		t.Error("Row was not inserted properly")
	}

}
