package tests

import (
	"fmt"
	"testing"

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

	// Create a new parser
	parser, err := parse.NewParser("../logs/conn.log", false)
	if err != nil {
		t.Fatal(err)
	}

	// Grab the specific fields to parse from config
	parser.SetFields(conf.Parser["conn"].Fields)

	// Grab the index of the specific values that will be parsed
	err = parser.GetIndexOfFields()
	if err != nil {
		t.Fatal(err)
	}

	// How many rows do you want to buffer
	parser.CreateBuffer(100)

	go parser.BufferRow(parse.ConnParse)

	err = db.InsertBatch(parser.Row, "conn", len(parser.Fields()))
	if err != nil {
		t.Fatal(err)
	}

}

func TestParseDns(t *testing.T) {

	// Create a new parser
	parser, err := parse.NewParser("../logs/dns.log", false)
	if err != nil {
		t.Fatal(err)
	}

	// Grab the specific fields to parse from config
	parser.SetFields(conf.Parser["dns"].Fields)

	// Grab the index of the specific values that will be parsed
	err = parser.GetIndexOfFields()
	if err != nil {
		t.Fatal(err)
	}

	// How many rows do you want to buffer
	parser.CreateBuffer(100)

	go parser.BufferRow(parse.DnsParse)

	// Insert into db
	err = db.InsertBatch(parser.Row, "dns", len(parser.Fields()))
	if err != nil {
		t.Fatal(err)
	}

	topDomains, err := db.TopFiveDomains()

	fmt.Println(topDomains)
	fmt.Println(err)

}
