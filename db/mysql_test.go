package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/amadeovezz/gobro/config"
	"github.com/stretchr/testify/assert"
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

func TestQueryTopDomains(t *testing.T) {
	assert := assert.New(t)

	// Create buffer
	dnsBuffer := make(chan []string, 20)

	record := []string{"1476454019", "axb3912eK345", "10.0.0.20", "22", "8.8.8.8",
		"88", "udp", "12", "reddit.com", "0", "0", "0", "-", "-",
		"-", "-", "-", "-", "-", "-", "-", "-", "-"}

	dnsBuffer <- record

	record = []string{"1476454017", "axb3912eK345", "10.0.0.20", "22", "8.8.8.8",
		"88", "udp", "12", "reddit.com", "0", "0", "0", "-", "-",
		"-", "-", "-", "-", "-", "-", "-", "-", "-"}

	dnsBuffer <- record

	record = []string{"1476454018", "axb3912eK345", "10.0.0.20", "22", "8.8.8.8",
		"88", "udp", "12", "reddit.com", "0", "0", "0", "-", "-",
		"-", "-", "-", "-", "-", "-", "-", "-", "-"}

	dnsBuffer <- record

	record = []string{"1476454019", "axb3912eK345", "10.0.0.20", "22", "8.8.8.8",
		"88", "udp", "12", "youtube.com", "0", "0", "0", "-", "-",
		"-", "-", "-", "-", "-", "-", "-", "-", "-"}

	dnsBuffer <- record

	record = []string{"1476454019", "axb3912eK345", "10.0.0.20", "22", "8.8.8.8",
		"88", "udp", "12", "youtube.com", "0", "0", "0", "-", "-",
		"-", "-", "-", "-", "-", "-", "-", "-", "-"}

	dnsBuffer <- record

	record = []string{"1476454019", "axb3912eK345", "10.0.0.20", "22", "8.8.8.8",
		"88", "udp", "12", "google.com", "0", "0", "0", "-", "-",
		"-", "-", "-", "-", "-", "-", "-", "-", "-"}

	dnsBuffer <- record

	record = []string{"1476454049", "axb3912eK345", "10.0.0.20", "22", "8.8.8.8",
		"88", "udp", "12", "google.com", "0", "0", "0", "-", "-",
		"-", "-", "-", "-", "-", "-", "-", "-", "-"}

	dnsBuffer <- record

	record = []string{"1476452019", "axb3912eK345", "10.0.0.20", "22", "8.8.8.8",
		"88", "udp", "12", "bro.org", "0", "0", "0", "-", "-",
		"-", "-", "-", "-", "-", "-", "-", "-", "-"}

	dnsBuffer <- record

	record = []string{"1436454019", "axb3912eK345", "10.0.0.20", "22", "8.8.8.8",
		"88", "udp", "12", "golang.org", "0", "0", "0", "-", "-",
		"-", "-", "-", "-", "-", "-", "-", "-", "-"}

	dnsBuffer <- record

	close(dnsBuffer)

	// Insert into db
	err := InsertBatch(dnsBuffer, "dns", len(record))
	if err != nil {
		t.Error(err)
	}

	domains, err := TopFiveDomains()

	assert.Equal(domains[0].Query, "reddit.com", "wrong domain as most visited")
	assert.Equal(domains[0].Count, 3, "wrong count as most visited")

	assert.Equal(domains[1].Query, "youtube.com", "wrong domain as second most visited")
	assert.Equal(domains[1].Count, 2, "wrong count as second most visited")

	assert.Equal(domains[2].Query, "google.com", "wrong domain as second most visited")
	assert.Equal(domains[2].Count, 2, "wrong count as second most visited")

	assert.Equal(domains[3].Query, "bro.org", "wrong domain as third most visited")
	assert.Equal(domains[3].Count, 1, "wrong count as third most visited")

	assert.Equal(domains[4].Query, "golang.org", "wrong domain as third most visited")
	assert.Equal(domains[4].Count, 1, "wrong count as third most visited")

}
