package app

import (
	"encoding/json"

	squaremicroservices "github.com/longvu727/FootballSquaresLibs/services/square_microservices"
	"github.com/longvu727/FootballSquaresLibs/util/resources"
)

type GetSquareParams struct {
	SquareID int `json:"square_id"`
}

type GetSquareResponse struct {
	squaremicroservices.Square
	ErrorMessage string `json:"error_message"`
}

func (response GetSquareResponse) ToJson() []byte {
	jsonStr, _ := json.Marshal(response)
	return jsonStr
}

func (squareApp *SquareApp) GetDBSquare(getSquareParams GetSquareParams, resources *resources.Resources) (*GetSquareResponse, error) {
	var getSquareResponse GetSquareResponse

	squareRow, err := resources.DB.GetSquare(resources.Context, int32(getSquareParams.SquareID))
	if err != nil {
		return &getSquareResponse, err
	}

	getSquareResponse.SquareID = int(squareRow.SquareID)
	getSquareResponse.SquareGUID = squareRow.SquareGuid
	getSquareResponse.SquareSize = int(squareRow.SquareSize.Int32)
	getSquareResponse.RowPoints = squareRow.RowPoints.String
	getSquareResponse.ColumnPoints = squareRow.ColumnPoints.String

	return &getSquareResponse, nil
}
