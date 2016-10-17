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
