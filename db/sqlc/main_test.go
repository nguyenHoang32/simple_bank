package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"simple_bank/util"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")

	if err != nil {
		log.Fatal("cannot load config", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(testDB)
	fmt.Println(testQueries)
	os.Exit(m.Run())
}
