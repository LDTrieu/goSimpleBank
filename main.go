package main

import (
	"database/sql"
	"log"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"

	"github.com/LDTrieu/goSimpleBank/api"
	db "github.com/LDTrieu/goSimpleBank/db/sqlc"
	"github.com/LDTrieu/goSimpleBank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
