package app

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"squaremicroservices/db"

	"github.com/google/uuid"
)

type CreateSquareParams struct {
	SquareSize int32  `json:"square_size"`
	TeamA      string `json:"team_a"`
	TeamB      string `json:"team_b"`
}
type CreateSquareResponse struct {
	SquareGUID   string `json:"square_guid"`
	ErrorMessage string `json:"error_message"`
}

func (response CreateSquareResponse) ToJson() []byte {
	jsonStr, _ := json.Marshal(response)
	return jsonStr
}

func CreateDBSquare(ctx context.Context, request *http.Request, dbConnect *db.MySQL) (CreateSquareResponse, error) {
	var createSquareResponse CreateSquareResponse
	var createSquareParams CreateSquareParams

	json.NewDecoder(request.Body).Decode(&createSquareParams)

	squareGuid := uuid.New()
	createSquareResponse.SquareGUID = squareGuid.String()

	_, err := dbConnect.QUERIES.CreateSquare(ctx, db.CreateSquareParams{
		SquareGuid: squareGuid.String(),
		SquareSize: sql.NullInt32{Int32: createSquareParams.SquareSize, Valid: true},
	})
	if err != nil {
		return createSquareResponse, err
	}

	gameGuid := uuid.New()
	_, err = dbConnect.QUERIES.CreateGames(ctx, db.CreateGamesParams{
		GameGuid: gameGuid.String(),
		Sport:    sql.NullString{String: "football", Valid: true},
		TeamA:    sql.NullString{String: createSquareParams.TeamA, Valid: true},
		TeamB:    sql.NullString{String: createSquareParams.TeamB, Valid: true},
	})

	return createSquareResponse, err
}
