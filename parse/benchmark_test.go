package parse

import (
	"log"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Logs logs
}

type logs struct {
	PathToLog      string
	PathToLargeLog string
}

var conf Config

func (c *Config) SetupConfig(path string) {
	filename, _ := filepath.Abs(path)
	if _, err := toml.DecodeFile(filename, c); err != nil {
		log.Fatal("setupConfig() could not decode toml, err: ", err)
	}
}

func TestMain(m *testing.M) {

	conf.SetupConfig("config.toml")

	m.Run()

}

func BenchmarkWithAutoInitialization(b *testing.B) {

	parser, err := NewParser(conf.Logs.PathToLog, true)
	if err != nil {
		log.Panic(err)
	}

	fields, err := parser.ParseAllFields()
	if err != nil {
		log.Panic(err)
	}

	parser.SetFields(fields)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := parser.AutoCreateBuffer()
		if err != nil {
			b.Fatal(err)
		}
		parser.BufferRow()
	}

}

func BenchmarkWithoutAutoInitialization(b *testing.B) {

	parser, err := NewParser(conf.Logs.PathToLog, true)
	if err != nil {
		log.Panic(err)
	}

	fields, err := parser.ParseAllFields()
	if err != nil {
		log.Panic(err)
	}

	parser.SetFields(fields)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parser.CreateBuffer(1468)
		parser.BufferRow()
	}

}

func BenchmarkLargeLogFile(b *testing.B) {

	parser, err := NewParser(conf.Logs.PathToLargeLog, true)
	if err != nil {
		log.Panic(err)
	}

	fields, err := parser.ParseAllFields()
	if err != nil {
		log.Panic(err)
	}

	parser.SetFields(fields)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := parser.AutoCreateBuffer()
		if err != nil {
			b.Fatal(err)
		}
		parser.BufferRow()
	}

}

func BenchmarkCountingLinesWithWC(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_, err := exec.Command("wc", "-l", "conn.log").Output()
		if err != nil {
			b.Fatal(err)
		}
	}

}

func BenchmarkCountingLinesWithGo(b *testing.B) {

	parser, err := NewParser(conf.Logs.PathToLog, true)
	if err != nil {
		log.Panic(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := parser.CountLines()
		if err != nil {
			b.Fatal(err)
		}
	}

}
