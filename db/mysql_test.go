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

func TestInsert(t *testing.T) {

	conn := createConn(1476454019, "axb3912eK345", "10.0.0.20", "22", "8.8.8.8",
		"88", "tcp", "facebook", 10.001, 200, 4000, 30, 400)

	conn2 := createConn(1476454020, "bxb3912rReK345", "10.0.0.20", "22", "8.8.8.8",
		"88", "udp", "", 10.001, 200, 4000, 30, 400)

	conn3 := createConn(1476854020, "o38912rReK345", "10.0.0.20", "22", "8.8.8.8",
		"88", "", "linkedin", 10.001, 200, 4000, 30, 400)

	conn4 := createConn(1476454020, "c89Ij12rReK345", "10.0.0.20", "22", "8.8.8.8",
		"88", "udp", "facebook", 10.001, 200, 4000, 30, 400)

	connArray := []Conn{conn, conn2, conn3, conn4}

	err := InsertBatchIntoConn(connArray)
	if err != nil {
		t.Error(err)
	}

}

func createConn(time uint64,
	connUID,
	origIp,
	origPort,
	respIp,
	respPort,
	proto,
	service string,
	duration float64,
	inBytes uint64,
	outBytes,
	inPackets,
	outPackets uint64) Conn {

	var conn Conn
	conn.time = time
	conn.connUID = connUID
	conn.origIp = origIp
	conn.origPort = origPort
	conn.respIp = respIp
	conn.respPort = respPort
	conn.proto = proto
	conn.service = service
	conn.duration = duration
	conn.inBytes = inBytes
	conn.outBytes = outBytes
	conn.inPackets = inPackets
	conn.outPackets = outPackets

	return conn

}
