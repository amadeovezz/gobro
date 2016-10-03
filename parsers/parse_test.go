package parsers

import (
	"testing"

	log "github.com/Sirupsen/logrus"
)

func TestGetFields(t *testing.T) {

	parser, err := NewParser("../logs/conn.log")
	if err != nil {
		log.Fatal(err)
	}

	err = parser.ParseFields()
	if err != nil {
		log.Fatal(err)
	}

	log.Info(parser.GetFields())

}
