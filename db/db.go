package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type DB struct {
    Client *sql.DB
}

func NewDBStorage(cfg mysql.Config) (*sql.DB, error) {
    db, err := sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    return db, nil
}

func NewDatabase(dsn string) (*DB, error) {
    dbConn := &DB{}

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    
    err = db.Ping()
    if err != nil {
        return nil, err
    }

    return dbConn, nil
}

func checkDB(d *sql.DB) error {
    err := d.Ping()
    if err != nil {
        fmt.Println("Error", err)
        return err
    }
    fmt.Println("*** Pinged database successfully ***")
    return nil
}
