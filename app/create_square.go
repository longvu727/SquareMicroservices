package app

import (
	"database/sql"
	"encoding/json"

	"github.com/longvu727/FootballSquaresLibs/DB/db"
	"github.com/longvu727/FootballSquaresLibs/util/resources"

	"github.com/google/uuid"
)

type CreateSquareParams struct {
	SquareSize int32 `json:"square_size"`
}
type CreateSquareResponse struct {
	SquareID     int64  `json:"square_id"`
	SquareGUID   string `json:"square_guid"`
	ErrorMessage string `json:"error_message"`
}

func (response CreateSquareResponse) ToJson() []byte {
	jsonStr, _ := json.Marshal(response)
	return jsonStr
}

func (squareApp *SquareApp) CreateDBSquare(createSquareParams CreateSquareParams, resources *resources.Resources) (*CreateSquareResponse, error) {
	var createSquareResponse CreateSquareResponse

	squareGuid := (uuid.New()).String()

	squareID, err := resources.DB.CreateSquare(resources.Context, db.CreateSquareParams{
		SquareGuid: squareGuid,
		SquareSize: sql.NullInt32{Int32: int32(createSquareParams.SquareSize), Valid: true},
	})
	if err != nil {
		return &createSquareResponse, err
	}

	createSquareResponse.SquareGUID = squareGuid
	createSquareResponse.SquareID = squareID

	return &createSquareResponse, nil
}
