package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbdriver string = "postgres"
	dbsource string = "postgres://root:secret@localhost:2345/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbdriver, dbsource)
	if err != nil {
		log.Fatal("cannot connect to db")
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
