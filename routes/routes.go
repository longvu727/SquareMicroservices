package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"squaremicroservices/app"
	"github.com/longvu727/FootballSquaresLibs/DB/db"
)

type Handler = func(writer http.ResponseWriter, request *http.Request)

func Register(db *db.MySQL, ctx context.Context) {
	log.Println("Registering routes")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		home(w, r)
	})

	http.HandleFunc(http.MethodPost+" /CreateSquare", func(w http.ResponseWriter, r *http.Request) {
		createSquare(w, r, db, ctx)
	})

	http.HandleFunc(http.MethodPost+" /GetSquare", func(w http.ResponseWriter, r *http.Request) {
		getSquare(w, r, db, ctx)
	})
}

func home(writer http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(writer, "{\"Acknowledged\": true}")
}

func createSquare(writer http.ResponseWriter, request *http.Request, dbConnect *db.MySQL, ctx context.Context) {
	log.Printf("Received request for %s\n", request.URL.Path)

	writer.Header().Set("Content-Type", "application/json")

	createSquareResponse, err := app.CreateDBSquare(ctx, request, dbConnect)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		createSquareResponse.ErrorMessage = `Unable to create square`
		writer.Write(createSquareResponse.ToJson())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(createSquareResponse.ToJson())
}

func getSquare(writer http.ResponseWriter, request *http.Request, dbConnect *db.MySQL, ctx context.Context) {
	log.Printf("Received request for %s\n", request.URL.Path)

	writer.Header().Set("Content-Type", "application/json")

	getSquareResponse, err := app.GetDBSquare(ctx, request, dbConnect)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		getSquareResponse.ErrorMessage = `Unable to get square`
		writer.Write(getSquareResponse.ToJson())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(getSquareResponse.ToJson())
}
