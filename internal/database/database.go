package database

import (
	"database/sql"
	"strings"

	"github.com/chiragsoni81245/net-sentinel/internal/config"
	_ "github.com/mattn/go-sqlite3"
)


func GetDB(config *config.Config) (*sql.DB, error) {
    var DB *sql.DB
	var err error

	// Connect to the database
    DB, err = sql.Open("sqlite3", strings.Replace(config.Database.URI, "sqlite3:/", ".", 1))
	if err != nil {
        return nil, err
	}

	// Ping the database to check if the connection is successful
	if err = DB.Ping(); err != nil {
        return nil, err
	}

    return DB, err
}
