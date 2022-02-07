package psql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	driverName = "postgres"
	host       = "localhost"
	port       = 5432
	user       = "postgres"
	password   = "home"
	dbname     = "test"
)

var dataSourceName = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

type myRepo struct {
	db *sql.DB
}

func NewRepo() (*myRepo, error) {
	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &myRepo{db}, nil
}
