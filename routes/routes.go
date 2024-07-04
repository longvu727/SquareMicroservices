package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"squaremicroservices/app"

	"github.com/longvu727/FootballSquaresLibs/util/resources"
)

type RoutesInterface interface {
	Register(resources *resources.Resources) *http.ServeMux
}

type Routes struct {
	Apps app.Square
}

type Handler = func(writer http.ResponseWriter, request *http.Request)

func NewRoutes() RoutesInterface {
	return &Routes{
		Apps: app.NewSquareApp(),
	}
}

func (routes *Routes) Register(resources *resources.Resources) *http.ServeMux {
	log.Println("Registering routes")
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		routes.home(w, r)
	})

	mux.HandleFunc(http.MethodPost+" /CreateSquare", func(w http.ResponseWriter, r *http.Request) {
		routes.createSquare(w, r, resources)
	})

	mux.HandleFunc(http.MethodPost+" /GetSquare", func(w http.ResponseWriter, r *http.Request) {
		routes.getSquare(w, r, resources)
	})

	return mux
}

func (routes *Routes) home(writer http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(writer, "{\"Acknowledged\": true}")
}

func (routes *Routes) createSquare(writer http.ResponseWriter, request *http.Request, resources *resources.Resources) {
	log.Printf("Received request for %s\n", request.URL.Path)

	writer.Header().Set("Content-Type", "application/json")

	var createSquareParams app.CreateSquareParams
	json.NewDecoder(request.Body).Decode(&createSquareParams)

	createSquareResponse, err := routes.Apps.CreateDBSquare(createSquareParams, resources)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		createSquareResponse.ErrorMessage = `Unable to create square`
		writer.Write(createSquareResponse.ToJson())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(createSquareResponse.ToJson())
}

func (routes *Routes) getSquare(writer http.ResponseWriter, request *http.Request, resources *resources.Resources) {
	log.Printf("Received request for %s\n", request.URL.Path)

	writer.Header().Set("Content-Type", "application/json")

	var getSquareParams app.GetSquareParams
	json.NewDecoder(request.Body).Decode(&getSquareParams)

	getSquareResponse, err := routes.Apps.GetDBSquare(getSquareParams, resources)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		getSquareResponse.ErrorMessage = `Unable to get square`
		writer.Write(getSquareResponse.ToJson())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(getSquareResponse.ToJson())
}
