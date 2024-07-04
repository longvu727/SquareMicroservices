package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"squaremicroservices/app"

	"github.com/longvu727/FootballSquaresLibs/util/resources"
)

type Handler = func(writer http.ResponseWriter, request *http.Request)

func Register(resources *resources.Resources) {
	log.Println("Registering routes")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		home(w, r)
	})

	http.HandleFunc(http.MethodPost+" /CreateSquare", func(w http.ResponseWriter, r *http.Request) {
		createSquare(w, r, resources)
	})

	http.HandleFunc(http.MethodPost+" /GetSquare", func(w http.ResponseWriter, r *http.Request) {
		getSquare(w, r, resources)
	})
}

func home(writer http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(writer, "{\"Acknowledged\": true}")
}

func createSquare(writer http.ResponseWriter, request *http.Request, resources *resources.Resources) {
	log.Printf("Received request for %s\n", request.URL.Path)

	writer.Header().Set("Content-Type", "application/json")

	var createSquareParams app.CreateSquareParams
	json.NewDecoder(request.Body).Decode(&createSquareParams)

	createSquareResponse, err := app.NewSquareApp().CreateDBSquare(createSquareParams, resources)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		createSquareResponse.ErrorMessage = `Unable to create square`
		writer.Write(createSquareResponse.ToJson())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(createSquareResponse.ToJson())
}

func getSquare(writer http.ResponseWriter, request *http.Request, resources *resources.Resources) {
	log.Printf("Received request for %s\n", request.URL.Path)

	writer.Header().Set("Content-Type", "application/json")

	var getSquareParams app.GetSquareParams
	json.NewDecoder(request.Body).Decode(&getSquareParams)

	getSquareResponse, err := app.NewSquareApp().GetDBSquare(getSquareParams, resources)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		getSquareResponse.ErrorMessage = `Unable to get square`
		writer.Write(getSquareResponse.ToJson())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(getSquareResponse.ToJson())
}
