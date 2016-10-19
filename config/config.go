// Package config contains all the file dependencies for gobro: config.toml
// and schema.sql. It also provides utility functions to access the toml
// objects defined in config.toml. It is important to note that when defining
// what fields to parse in config.toml, the same fields must be included
// in schema.sql. However, all fields with ".", must be replaced in
// the with "_" in schema.sql.
package config

import (
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Title  string
	Parser map[string]parser
}

type parser struct {
	Fields []string
}

func (c *Config) SetupConfig(path string) {
	filename, _ := filepath.Abs(path)
	if _, err := toml.DecodeFile(filename, c); err != nil {
		log.Fatal("setupConfig() could not decode toml, err: ", err)
	}
}
