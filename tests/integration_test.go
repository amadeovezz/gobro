package tests

import (
	"errors"
	"fmt"
	"testing"

	"golang.org/x/net/publicsuffix"

	"github.com/amadeovezz/gobro/config"
	"github.com/amadeovezz/gobro/db"
	"github.com/amadeovezz/gobro/parse"
)

var conf config.Config

func TestMain(m *testing.M) {

	conf.SetupConfig("config.toml")

	err := db.InitDB(
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

func TestParseConn(t *testing.T) {

	// Create a new parser with specific fields, and raw entries
	parser, err := parse.NewParser("../logs/conn.log", false, true)
	if err != nil {
		t.Fatal(err)
	}

	// Grab the specific fields to parse from config
	parser.SetFields(conf.Parser["conn"].Fields)

	// How many rows do you want to buffer
	parser.CreateBuffer(100)

	// No need for further data augmentation
	go parser.BufferRow(func(fields, row []string) ([]string, error) {
		return row, nil
	})

	err = db.InsertBatch(parser.Row, "conn", len(parser.Fields()))
	if err != nil {
		t.Fatal(err)
	}

}

func DnsParse(fields, row []string) ([]string, error) {

	var newRow []string
	newRow = row

	for i, field := range fields {

		if field == "query" {
			if row[i] == "-" {
				// Maybe change this to bool
				return nil, errors.New("No query information provided")
			}
			secondLevelDomain, err := publicsuffix.EffectiveTLDPlusOne(newRow[i])
			if err == nil {
				newRow[i] = secondLevelDomain
			}
		}

	}

	return newRow, nil

}

func TestParseDns(t *testing.T) {

	// Create a new parser with specific field and augemented entries
	parser, err := parse.NewParser("../logs/dns.log", false, false)
	if err != nil {
		t.Fatal(err)
	}

	// Grab the specific fields to parse from config
	parser.SetFields(conf.Parser["dns"].Fields)

	// How many rows do you want to buffer
	parser.CreateBuffer(100)

	// Strip out uncessary domain information
	go parser.BufferRow(DnsParse)

	// Insert into db
	err = db.InsertBatch(parser.Row, "dns", len(parser.Fields()))
	if err != nil {
		t.Fatal(err)
	}

	topDomains, err := db.TopFiveDomains()

	fmt.Println(topDomains)
	fmt.Println(err)

}
