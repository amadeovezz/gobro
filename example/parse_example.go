package example

import (
	"fmt"

	"github.com/amadeovezz/gobro/config"
	"github.com/amadeovezz/gobro/parse"
)

func ParseWithConfigFields() {

	var config config.Config
	config.SetupConfig("config/config.toml")

	// Create a new parser and grab raw entries
	parser, err := parse.NewParser("/logs/example.log", true)
	if err != nil {
		t.Error(err)
	}

	// Grab the specific fields to parse from config
	parser.SetFields(config.Parser["conn"].Fields)

	// Grab the index of the specific values that will be parsed
	err = parser.GrabIndexOfFields()
	if err != nil {
		t.Error(err)
	}

	// How many rows do you want to buffer
	parser.CreateBuffer(10)

	// The definition of ParseConn can just be empty, but it still needs
	// to be declared
	go parser.BufferRow(parser.ParseConn)

	// Do stuff with data
	for row := range parser.Row {
		fmt.Println(row)
	}

}

func ParseAllFields() {

	// Create a new parser and manipulate specific entries
	parser, err := NewParser("../logs/conn.log", false)
	if err != nil {
		t.Error(err)
	}

	// Grab all the fields in the bro log
	fields, err := parser.ParseAllFields()
	if err != nil {
		t.Error(err)
	}

	// Assign the fields you want to parse
	parser.SetFields(fields)

	// How many rows do you want to buffer
	parser.CreateBuffer(10)

	// Define and write a ParseConn to manipulate the data
	go parser.BufferRow(parser.ParseConn)

	// Do stuff with data
	for row := range parser.Row {
		fmt.Println(row)
	}

}
