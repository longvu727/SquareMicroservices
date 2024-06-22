package main

import (
	"fmt"
	"log"
	"net/http"
	"squaremicroservices/routes"
	"squaremicroservices/util"
)

func main() {
	config, err := util.LoadConfig(".", "app", "env")
	if err != nil {
		log.Fatal(err)
	}

	handler(fmt.Sprintf(":%s", config.PORT))
}

func handler(address string) error {
	routes.Register()
	log.Printf("Listening on %s", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
