# gobro

gobro is a extendable golang toolkit to work with BRO IDS data.

# Features for version 0.1.0

* Expose a generic parsing api 
* Expose a generic api to populate a db using golangs mysql package
* Implement a full unit test suite for /parsers and /db 
* Implement integration tests to test end to end data flow
* Provide documentation on gobro, and use godoc

gobro uses semantic versioning 

# Project structure 

packages:

/parsers: logic for parsing BRO logs 

/db: db inserts and general utility functions 

/tests: for integration tests 

/config: includes toml file and schema for launching gobro

/doc: information about the type of statistics that are reported on

# Dependencies

* toml config file located in /config
* mysql schema (optional) also located in /config

# Additional links

https://www.bro.org/
http://gauss.ececs.uc.edu/Courses/c6055/pdf/bro_log_vars.pdf

# Language

### Bro logs

* separator: how bro logs are delimited
* fields: column names in bro logs
* entry: a single line in a bro log (the actual data)

# Example: Parsing bro logs 

```go
package main

import (
	"fmt"
	"log"

	"github.com/amadeovezz/gobro/config"
	"github.com/amadeovezz/gobro/parse"
)

func main() {

	// Config settings
	var conf config.Config
	conf.SetupConfig("config.toml")

	// Create a new parser with specific fields from config and parse raw entries
	parser, err := parse.NewParser("conn.log", false, true)
	if err != nil {
		log.Panic(err)
	}

	// Grab the specific fields to parse from config
	parser.SetFields(conf.Parser["conn"].Fields)

	// How many rows do you want to buffer
	parser.CreateBuffer(100)

	// If we don't want to modify the fields further
	go parser.BufferRow(func(fields, row []string) ([]string, error) {
		return row, nil
	})

	for data := range parser.Row {
		fmt.Println(data)
	}

}
```

# Example: Parsing a bro log and writing the data to mysql 

```go
package main

import (
	"errors"
	"log"

	"golang.org/x/net/publicsuffix"

	"github.com/amadeovezz/gobro/config"
	"github.com/amadeovezz/gobro/db"
	"github.com/amadeovezz/gobro/parse"
)

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

func main() {

	// Config settings
	var conf config.Config
	conf.SetupConfig("config.toml")
	err := db.InitDB(
		conf.DB.Username,
		conf.DB.Password,
		conf.DB.IP,
		conf.DB.Port,
		conf.DB.DatabaseName,
	)

	// Create a new parser with specific fields from config and augement the raw entries
	parser, err := parse.NewParser("dns.log", false, false)
	if err != nil {
		log.Panic(err)
	}

	// Grab the specific fields to parse from config
	parser.SetFields(conf.Parser["dns"].Fields)

	// How many rows do you want to buffer
	parser.CreateBuffer(100)

	// Lets manipulate the dns "query" field 
	go parser.BufferRow(DnsParse)

	// Insert all rows into db
	err = db.InsertBatch(parser.Row, "conn", len(parser.Fields()))
	if err != nil {
		log.Panic(err)
	}

}
```

# Note

This project is still in the planning stage. Goals, releases, project structure, architecture, etc... are all subject to change


