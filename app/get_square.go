package app

import (
	"context"
	"encoding/json"
	"net/http"
	"squaremicroservices/db"
)

type GetSquareParams struct {
	SquareSize int32  `json:"square_size"`
	TeamA      string `json:"team_a"`
	TeamB      string `json:"team_b"`
}
type GetSquareResponse struct {
	SquareGUID   string `json:"square_guid"`
	ErrorMessage string `json:"error_message"`
}

func (response GetSquareResponse) ToJson() []byte {
	jsonStr, _ := json.Marshal(response)
	return jsonStr
}

func GetDBSquare(ctx context.Context, request *http.Request, dbConnect *db.MySQL) (*GetSquareResponse, error) {
	return nil, nil
}
