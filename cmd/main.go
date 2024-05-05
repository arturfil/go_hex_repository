package main

import (
	"database/sql"
	"log"

	"github.com/arturfil/go_repository_hex/cmd/api"
	"github.com/arturfil/go_repository_hex/config"
	"github.com/arturfil/go_repository_hex/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewDBStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

    initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
    err := db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("*** db successfully connected ***")
}
