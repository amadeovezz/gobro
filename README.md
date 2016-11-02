# gobro

gobro is an extendable golang toolkit to work with BRO IDS data.

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
* mysql schema also located in /config

A note about the toml and the schema file:

config.toml: when defining fields in the [parser] section of the
toml file, make sure to name the fields exactly as they appear 
in the bro log. 

schmema.sql: when defining the tables in the schema, make
sure to name each column exactly how the fields are named
in config.toml. Except all "." must be replaced with an "_". 

# Additional links

https://www.bro.org/
http://gauss.ececs.uc.edu/Courses/c6055/pdf/bro_log_vars.pdf

# Language

### Bro logs

* separator: how bro logs are delimited
* fields: column names in bro logs
* entry: a single line in a bro log (the actual data)


# Example 1 : Parsing bro logs

In this example we are creating a parser and storing the values of a 
bro log in memory. The boolean parameter "false" for NewParser() indicates 
that we are going to be parsing specific fields instead of all the fields
in the bro log. 

```go
package main

import (
	"fmt"
	"log"

	"github.com/amadeovezz/gobro/config"
	"github.com/amadeovezz/gobro/parse"
)

func main() {

	var conf config.Config
	conf.SetupConfig("config.toml")
	
	parser, err := parse.NewParser("conn.log", false)
	if err != nil {
		log.Panic(err)
	}

	parser.SetFields(conf.Parser["conn"].Fields)

	parser.CreateBuffer(100)

	go parser.BufferRow()

	for data := range parser.Row {
		fmt.Println(data)
	}

}
```

# Example 2 : Parsing with entry manipulations and storing data in sql

In this example, we pass in true to the boolean argument of
NewParser. ParseAllFields() must be called and then the returned fields must
be passed into SetFields(). In addition we have 
decided to modify and augment a certain field in the bro log. 
We create a function DnsParse() and pass it into BufferRow to 
perform additional data manipulations on the raw entries. gobro 
also comes built in with the option to store the data in SQL databases. 

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

	var conf config.Config
	conf.SetupConfig("config.toml")
	err := db.InitDB(
		conf.DB.Username,
		conf.DB.Password,
		conf.DB.IP,
		conf.DB.Port,
		conf.DB.DatabaseName,
	)

	parser, err := parse.NewParser("dns.log", true)
	if err != nil {
		log.Panic(err)
	}
	
	fields, err := parser.ParseAllFields()
	if err != nil {
		log.Panic(err)
	}

	parser.SetFields(fields)

	parser.CreateBuffer(100)

	go parser.BufferRow(DnsParse)

	err = db.InsertBatch(parser.Row, "conn", len(parser.Fields()))
	if err != nil {
		log.Panic(err)
	}

}
```



# Note

This project is still in the planning stage. Goals, releases, project structure, architecture, etc... are all subject to change


