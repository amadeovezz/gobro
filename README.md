# gobro

gobro is an extendable golang toolkit to work with Bro IDS data.

# Features for version 0.1.0

* Expose a generic parsing api 
* Expose a the ability to populate a db using golangs mysql package
* Implement a full unit test suite for /parsers and /db 
* Implement integration tests to test end to end data flow
* Provide documentation on gobro, and use godoc

gobro uses semantic versioning 

# Project structure 

packages:

/parsers: logic for parsing Bro logs 

/db: db inserts and general utility functions 

/tests: for integration tests 

/config: includes toml file and schema for launching gobro

# Dependencies

* toml config file located in /config
* mysql schema also located in /config

A note about the toml and the schema file:

config.toml: when defining fields in the [parser] section of the
toml file, make sure to name the fields exactly as they appear 
in the Bro log. 

schema.sql: when defining the tables in the schema, make
sure to name each column exactly how the fields are named
in config.toml. Except all "." must be replaced with an "_". 

# Additional links

https://www.bro.org/

http://gauss.ececs.uc.edu/Courses/c6055/pdf/bro_log_vars.pdf

# Language

### Bro logs

* separator: how Bro logs are delimited
* fields: column names in Bro logs
* entry: a single line in a Bro log (the actual data)

# Testing 

All tests and benchmarks can be run via the test.py python script.

Benchmarks that are run from the /parser package require log files. You
can specify their path in /parser/config.toml.

Tests and benchmarks that are run from the /testing package require log files.
You can specify their path in /parser/config.toml.

All tests that require use of a db are transient.
That is, you do not have to install any db, as long as you have docker and 
docker-compose installed. The tests will spin up and down the db images automatically.

Run: python ./test.py -h for details

# Benchmarks

Below are some relevant benchmarks run from the parsing package:

The first two benchmarks were run with a conn.log file with 1468 lines (175K).
The last benchmark was run with a conn.log file with 100,000 lines (13M).

```
BenchmarkWithAutoInitialization-4      	    1000	   1938466 ns/op	  746983 B/op	    2936 allocs/op
BenchmarkWithoutAutoInitialization-4   	    1000	   1922132 ns/op	  746925 B/op	    2933 allocs/op
BenchmarkLargeLogFile-4                	      10	 157950895 ns/op	51978650 B/op	  200001 allocs/op
```

# Example 1 : Parsing bro logs

In this example we are creating a parser and storing the values of a 
Bro log in memory. The boolean parameter "false" for NewParser() indicates 
that we are going to be parsing specific fields instead of all the fields
in the Bro log. Note that we can also initialize the buffered channel
manually with a call to CreateBuffer(100). AutoCreateBuffer() counts the number
of lines in a log file, and creates a channel with that size. 
You can then access the values by ranging over the parser.Row
channel.


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

	parser.AutoCreateBuffer()
	//parser.CreateBuffer(100)

	go parser.BufferRow()

	for data := range parser.Row {
		fmt.Println(data)
	}

}
```

# Example 2 : Parsing with entry manipulations and storing data in sql

In this example, we decide to parse all fields of the bro log, and thus we pass 
in true as the boolean argument of NewParser. 
ParseAllFields() must be called and then the returned fields must
be passed into SetFields(). In addition we have 
decided to modify and augment a certain field in the Bro log. 
We create a function DnsParse() and pass it into BufferRow to 
perform additional data manipulations on the raw entries. gobro 
also comes built in with the option to store the data in SQL databases, which
is demonstrated by the call to InsertBatch().

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

	parser.AutoCreateBuffer()

	go parser.BufferRow(DnsParse)

	err = db.InsertBatch(parser.Row, "conn", len(parser.Fields()))
	if err != nil {
		log.Panic(err)
	}

}
```


# Note

This project is still in the planning stage. Goals, releases, project structure, architecture, etc... are all subject to change
