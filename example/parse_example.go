package example

import (
	"fmt"

	"github.com/amadeovezz/gobro/config"
	"github.com/amadeovezz/gobro/parse"
)

func ParseWithConfigFields() {

	var config config.Config
	config.SetupConfig("config/config.toml")

	// Create a new parser
	parser, err := parse.NewParser("/logs/example.log")
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

	go parser.BufferRow()

	// Do stuff with data
	for row := range parser.Row {
		fmt.Println(row)
	}

}

func ParseAllFields() {

	// Create a new parser
	parser, err := NewParser("../logs/conn.log")
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

	go parser.BufferRow()

	// Do stuff with data
	for row := range parser.Row {
		fmt.Println(row)
	}

}
