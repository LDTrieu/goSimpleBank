package main

import (
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

func main() {
	fmt.Println("Test")
	/*
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
	*/

}
