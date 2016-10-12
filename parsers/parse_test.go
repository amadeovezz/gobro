package parsers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldsParsing(t *testing.T) {
	assert := assert.New(t)

	// Note this conn log does not have all the default values
	// it would normally contain. It is shortened for testing purposes
	parser, err := NewParser("../logs/conn.log")
	if err != nil {
		t.Error(err)
	}

	err = parser.ParseFields()
	if err != nil {
		t.Error(err)
	}

	fields := parser.Fields()

	assert.Equal(fields[0], "ts", "parsed fields incorrectly")
	assert.Equal(fields[1], "uid", "parsed fields incorrectly")
	assert.Equal(fields[2], "id.orig_h", "parsed fields incorrectly")
	assert.Equal(fields[3], "id.orig_p", "parsed fields incorrectly")
	assert.Equal(fields[4], "id.resp_h", "parsed fields incorrectly")
	assert.Equal(fields[5], "id.resp_p", "parsed fields incorrectly")
	assert.Equal(fields[6], "proto", "parsed fields incorrectly")

}

func TestBufferRow(t *testing.T) {
	assert := assert.New(t)

	parser, err := NewParser("../logs/conn.log")
	if err != nil {
		t.Error(err)
	}

	err = parser.ParseFields()
	if err != nil {
		t.Error(err)
	}

	parser.CreateBuffer(10)

	go parser.BufferRow()

	for row := range parser.Row {
		assert.Equal(row["ts"], "1452684903.908400", "parsed entries incorrectly")
		assert.Equal(row["uid"], "CbOiIv2wbbH7F25W21", "parsed entries incorrectly")
		assert.Equal(row["id.orig_h"], "10.1.20.227", "parsed entries incorrectly")
		assert.Equal(row["id.orig_p"], "37218", "parsed entries incorrectly")
		assert.Equal(row["id.resp_h"], "204.238.149.187", "parsed entries incorrectly")
		assert.Equal(row["id.resp_p"], "443", "parsed entries incorrectly")
		assert.Equal(row["proto"], "tcp", "parsed entries incorrectly")
	}

}
