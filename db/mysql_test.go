package db

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"

	"github.com/BurntSushi/toml"
)

var conf Config

type Config struct {
	DB databaseConfig `toml:"database"`
}

type databaseConfig struct {
	Username     string
	Password     string
	IP           string
	Port         string
	DatabaseName string
}

func setUpConfig(config *Config, path string) {

	filename, _ := filepath.Abs(path)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panic(err)
	}

	if _, err = toml.Decode(string(data), config); err != nil {
		log.Panic(err)
	}
}

func TestMain(m *testing.M) {
	setUpConfig(&conf, "config.toml")

	err := InitDB(
		conf.DB.Username,
		conf.DB.Password,
		conf.DB.IP,
		conf.DB.Port,
		conf.DB.DatabaseName,
	)

	if err != nil {
		fmt.Println(err)
		return
	}
	m.Run()

}
