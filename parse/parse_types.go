package parse

import (
	"errors"

	"golang.org/x/net/publicsuffix"
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

func ConnParse(fields, row []string) ([]string, error) {

	return row, nil
}

func SSHParse(fields, row []string) ([]string, error) {

	return row, nil

}
