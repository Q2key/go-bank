package sqlc

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

var (
	testQueries *Queries
	testDb      *sql.DB
)

const (
	driver   = "postgres"
	dbSource = "postgres://postgres:postgres@localhost:5439/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error

	testDb, err = sql.Open(driver, dbSource)
	if err != nil {
		log.Fatalf(err.Error())
	}

	testQueries = &Queries{
		db: testDb,
	}

	m.Run()
}
