package parsers

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

// Each row in a bro log is represented by a map. The fields as keys
// and the rows as its corresponding values

type Parser struct {
	fields   []string
	filepath string
	Row      chan map[string]string
}

func NewParser(path string) (*Parser, error) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.New("File path does not exist")
	}

	p := new(Parser)
	p.filepath = path

	return p, nil

}

func (b *Parser) GetFields() []string {
	return b.fields
}

// TODO remove hardcoding of the seperator, slight chance it could be something
// other than tabs

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

func (p *Parser) BufferRow() {

	// Shouldn't have to check for errors here
	file, _ := os.Open(p.filepath)
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

			// Populate a row map
			rowMap := make(map[string]string)

			for i, field := range p.fields {

				if entry[i] == "-" {
					rowMap[field] = "N/A"
				} else {
					rowMap[field] = entry[i]
				}

			}

			// Add the row map to the buffer
			p.Row <- rowMap

		}

	}

	close(p.Row)
}
