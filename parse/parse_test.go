package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingAllFields(t *testing.T) {
	assert := assert.New(t)

	parser, err := NewParser("../logs/example.log", true)
	if err != nil {
		t.Fatal(err)
	}

	fields, err := parser.ParseAllFields()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(fields[0], "ts", "parsed fields incorrectly")
	assert.Equal(fields[1], "uid", "parsed fields incorrectly")
	assert.Equal(fields[2], "id.orig_h", "parsed fields incorrectly")
	assert.Equal(fields[3], "id.orig_p", "parsed fields incorrectly")
	assert.Equal(fields[4], "id.resp_h", "parsed fields incorrectly")
	assert.Equal(fields[5], "id.resp_p", "parsed fields incorrectly")
	assert.Equal(fields[6], "proto", "parsed fields incorrectly")

}

func TestBufferSpecificEntries(t *testing.T) {
	assert := assert.New(t)

	// Create a new parser with specific field and raw values
	parser, err := NewParser("../logs/example.log", false)
	if err != nil {
		t.Fatal(err)
	}

	fieldsToParse := []string{"ts", "id.orig_h", "id.resp_p", "proto"}

	parser.SetFields(fieldsToParse)

	parser.CreateBuffer(10)

	go parser.BufferRow()

	for row := range parser.Row {
		assert.Equal(row[0], "1452684903.908400", "parsed entries incorrectly")
		assert.Equal(row[1], "10.1.20.227", "parsed entries incorrectly")
		assert.Equal(row[2], "443", "parsed entries incorrectly")
		assert.Equal(row[3], "tcp", "parsed entries incorrectly")
	}

}

func TestBufferAllEntries(t *testing.T) {
	assert := assert.New(t)

	// Create a new parser with all fields from bro log and raw values
	parser, err := NewParser("../logs/example.log", true)
	if err != nil {
		t.Fatal(err)
	}

	fields, err := parser.ParseAllFields()
	if err != nil {
		t.Fatal(err)
	}

	parser.SetFields(fields)

	parser.CreateBuffer(10)

	go parser.BufferRow()

	for row := range parser.Row {
		assert.Equal(row[0], "1452684903.908400", "parsed entries incorrectly")
		assert.Equal(row[1], "CbOiIv2wbbH7F25W21", "parsed entries incorrectly")
		assert.Equal(row[2], "10.1.20.227", "parsed entries incorrectly")
		assert.Equal(row[3], "37218", "parsed entries incorrectly")
		assert.Equal(row[4], "204.238.149.187", "parsed entries incorrectly")
		assert.Equal(row[5], "443", "parsed entries incorrectly")
		assert.Equal(row[6], "tcp", "parsed entries incorrectly")
	}

}

func TestFieldsToUnderscore(t *testing.T) {

	assert := assert.New(t)

	parser, err := NewParser("../logs/example.log", true)
	if err != nil {
		t.Fatal(err)
	}

	fields, err := parser.ParseAllFields()
	if err != nil {
		t.Fatal(err)
	}

	parser.SetFields(fields)

	underScoreFields, err := parser.FieldsToUnderscore()

	assert.Equal(underScoreFields[2], "id_orig_h", "parsed fields incorrectly")
	assert.Equal(underScoreFields[3], "id_orig_p", "parsed fields incorrectly")
	assert.Equal(underScoreFields[4], "id_resp_h", "parsed fields incorrectly")
	assert.Equal(underScoreFields[5], "id_resp_p", "parsed fields incorrectly")
}
