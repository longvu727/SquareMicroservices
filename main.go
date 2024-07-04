package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"squaremicroservices/routes"

	"github.com/longvu727/FootballSquaresLibs/DB/db"
	"github.com/longvu727/FootballSquaresLibs/util"
	"github.com/longvu727/FootballSquaresLibs/util/resources"
)

func main() {
	config, err := util.LoadConfig("./env", "app", "env")
	if err != nil {
		log.Fatal(err)
	}

	mysql, err := db.NewMySQL(config)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	resources := resources.NewResources(config, mysql, ctx)
	handler(resources)
}

func handler(resources *resources.Resources) error {
	routes.Register(resources)

	address := fmt.Sprintf(":%s", resources.Config.PORT)
	log.Printf("Listening on %s", address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
