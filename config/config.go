/*
Package config provides utility functions to access the toml
objects defined in config.toml (examples of a config file is located at:
https://github.com/amadeovezz/gobro/tree/master/config)
It is important to note that when defining
what fields to parse in config.toml, the same fields must be included
in schema.sql. However, all fields with ".", must be replaced in
the with "_" in schema.sql.
*/
package config

import (
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Config represents config options for fields to parse and db options
type Config struct {
	Title  string
	Parser map[string]parser
	DB     database `toml:"database"`
}

type parser struct {
	Fields []string
}

type database struct {
	Username     string
	Password     string
	IP           string
	Port         string
	DatabaseName string
}

// SetupConfig parses a toml file and log.Fatal's if it cannot parse it correctly
func (c *Config) SetupConfig(path string) {
	filename, _ := filepath.Abs(path)
	if _, err := toml.DecodeFile(filename, c); err != nil {
		log.Fatal("SetupConfig() could not decode toml, err: ", err)
	}
}
