package app

import (
	"context"
	"encoding/json"
	"net/http"
	"squaremicroservices/db"
)

type GetSquareParams struct {
	SquareID int `json:"square_id"`
}

type GetSquareResponse struct {
	SquareID     int    `json:"square_id"`
	SquareGUID   string `json:"square_guid"`
	SquareSize   int    `json:"square_size"`
	RowPoints    string `json:"row_points"`
	ColumnPoints string `json:"column_points"`
	ErrorMessage string `json:"error_message"`
}

func (response GetSquareResponse) ToJson() []byte {
	jsonStr, _ := json.Marshal(response)
	return jsonStr
}

func GetDBSquare(ctx context.Context, request *http.Request, dbConnect *db.MySQL) (*GetSquareResponse, error) {
	var getSquareResponse GetSquareResponse
	var getSquareParams GetSquareParams
	json.NewDecoder(request.Body).Decode(&getSquareParams)

	squareRow, err := dbConnect.QUERIES.GetSquare(ctx, int32(getSquareParams.SquareID))
	if err != nil {
		return &getSquareResponse, err
	}

	getSquareResponse.SquareID = int(squareRow.SquareID)
	getSquareResponse.SquareGuid = squareRow.SquareGuid
	getSquareResponse.SquareSize = int(squareRow.SquareSize.Int32)
	getSquareResponse.RowPoints = squareRow.RowPoints.String
	getSquareResponse.ColumnPoints = squareRow.ColumnPoints.String

	return &getSquareResponse, nil
}
