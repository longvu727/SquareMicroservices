package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"squaremicroservices/db"
	"squaremicroservices/routes"
	"squaremicroservices/util"
)

func main() {
	config, err := util.LoadConfig("./env", "app", "env")
	log.SetOutput(os.Stdout)

	if err != nil {
		log.Fatal(err)
	}

	mysql, err := db.NewMySQL(config)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	handler(config, mysql, ctx)
}

func handler(config util.Config, db *db.MySQL, ctx context.Context) error {

	routes.Register(db, ctx)

	address := fmt.Sprintf(":%s", config.PORT)
	log.Printf("Listening on %s", address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
