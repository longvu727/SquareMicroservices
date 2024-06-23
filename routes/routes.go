package routes

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"squaremicroservices/db"

	"github.com/google/uuid"
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
}

func home(writer http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(writer, "{\"Acknowledged\": true}")
}

type CreateSquareParams struct {
	SizeLength int
}
type CreateSquareResponse struct {
	SquareGuid   string `json:"square_uuid"`
	ErrorMessage string `json:"error_message"`
}

func (response CreateSquareResponse) toJson() []byte {
	jsonStr, _ := json.Marshal(response)
	return jsonStr
}

func createSquare(writer http.ResponseWriter, request *http.Request, dbConnect *db.MySQL, ctx context.Context) {
	log.Printf("Received request for %s\n", request.URL.Path)

	var createSquareResponse CreateSquareResponse
	var createSquareParams CreateSquareParams

	json.NewDecoder(request.Body).Decode(&createSquareParams)

	squareGuid := uuid.New()
	createSquareResponse.SquareGuid = squareGuid.String()

	_, err := dbConnect.QUERIES.CreateSquare(ctx, db.CreateSquareParams{
		SquareGuid: squareGuid.String(),
		SideLength: sql.NullInt32{Int32: 10, Valid: true},
	})

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		createSquareResponse.SquareGuid = ""
		createSquareResponse.ErrorMessage = `Unable to create square`
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(createSquareResponse.toJson())
}
