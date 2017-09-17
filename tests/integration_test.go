package tests

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"testing"

	"golang.org/x/net/publicsuffix"

	"github.com/BurntSushi/toml"
	"github.com/amadeovezz/gobro/db"
	"github.com/amadeovezz/gobro/parse"
)

var conf Config

type Config struct {
	Logs   logs
	Parser map[string]parser
	DB     database `toml:"database"`
}

type parser struct {
	Fields []string
}

type database struct {
	Username     string
	Password     string
	IP           string
	Port         string
	DatabaseName string
}

type logs struct {
	PathToConn string
	PathToDns  string
}

func (c *Config) SetupConfig(path string) {
	filename, _ := filepath.Abs(path)
	if _, err := toml.DecodeFile(filename, c); err != nil {
		log.Fatal("setupConfig() could not decode toml, err: ", err)
	}
}

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
	parser, err := parse.NewParser(conf.Logs.PathToConn, false)
	if err != nil {
		t.Fatal(err)
	}

	// Grab the specific fields to parse from config
	parser.SetFields(conf.Parser["conn"].Fields)

	// How many rows do you want to buffer
	parser.CreateBuffer(100)

	go parser.BufferRow()

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
	parser, err := parse.NewParser(conf.Logs.PathToDns, false)
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
}

func BenchmarkParsingAndInsertingDb(b *testing.B) {

	for i := 0; i < b.N; i++ {

		parser, err := parse.NewParser(conf.Logs.PathToDns, false)
		if err != nil {
			b.Fatal(err)
		}

		parser.SetFields(conf.Parser["dns"].Fields)

		parser.AutoCreateBuffer()
		parser.BufferRow()

		// Insert into db
		err = db.InsertBatch(parser.Row, "dns", len(parser.Fields()))
		if err != nil {
			b.Fatal(err)
		}
	}
}
