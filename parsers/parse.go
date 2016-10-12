package parsers

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Parser manages the structure of a Bro log.
// Each row in a bro log is represented by a map. The fields as keys
// and the rows as its corresponding values.
// The rows are stored in a buffered channel.
type Parser struct {
	fields   []string
	filepath string
	Row      chan map[string]string
}

// NewParser validates the bro log exists and returns a new parser
// to perform parsing actions on.
func NewParser(path string) (*Parser, error) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.New("File path does not exist")
	}

	p := new(Parser)
	p.filepath = path

	return p, nil

}

// Fields returns the fields of a bro log.
func (b *Parser) Fields() []string {
	return b.fields
}

// TODO remove hardcoding of the seperator, it could be something
// other than tabs (research this)?

// ParseFields parses the fields of a bro log, and stores them in a
// slice. Their positions in the bro log correspond to their index's
// in the slice.
func (p *Parser) ParseFields() error {

	file, fileErr := os.Open(p.filepath)
	if fileErr != nil {
		return fileErr
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line[0:7] == "#fields" {

			// Lets make sure the fields row is not malformed
			if line[8:] == "" {
				return errors.New("Fields row is malformed")
			}

			p.fields = strings.Split(line[8:], "\t")
			break
		}

	}

	return nil

}

// CreateBuffer initializes the buffer. Without initialization, the channel
// will block on read's.
func (p *Parser) CreateBuffer(bufferSize int) {
	p.Row = make(chan map[string]string, bufferSize)
}

// BufferRow parses throught the entries (data) of a bro log, and maps
// them to the fields of the bro log. Every row is then pushed into the
// channel p.Row.
func (p *Parser) BufferRow() {

	if p.Row == nil {
		fmt.Println("Initialize nil channel, via CreateBuffer()")
		return
	}

	file, fileErr := os.Open(p.filepath)
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Any line without a # is a row with values
		if string(line[0]) != "#" {

			// Lets make sure the value row is not malformed
			if line[1:] == "" {
				continue
			}

			entry := strings.Split(line, "\t")

			// Skip this line if columns and values don't match
			if len(p.fields) != len(entry) {
				continue
			}

			// Create row
			rowMap := make(map[string]string)

			for i, field := range p.fields {

				if entry[i] == "-" {
					rowMap[field] = "N/A"
				} else {
					rowMap[field] = entry[i]
				}

			}

			p.Row <- rowMap

		}

	}

	close(p.Row)

}
