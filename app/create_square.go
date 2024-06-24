package app

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"squaremicroservices/db"

	"github.com/google/uuid"
)

type CreateSquareParams struct {
	Sport      string `json:"sport"`
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

func CreateDBSquare(ctx context.Context, request *http.Request, dbConnect *db.MySQL) (*CreateSquareResponse, error) {
	var createSquareParams CreateSquareParams
	json.NewDecoder(request.Body).Decode(&createSquareParams)

	if createSquareParams.Sport == "football" {
		return createDBFootballSquareGame(ctx, createSquareParams, dbConnect)
	}

	return nil, errors.New(`unknown sport`)
}

func createDBFootballSquareGame(ctx context.Context, createSquareParams CreateSquareParams, dbConnect *db.MySQL) (*CreateSquareResponse, error) {
	var createSquareResponse CreateSquareResponse
	squareGuid := uuid.New()

	squareID, err := insertSquare(ctx, dbConnect, squareGuid, int(createSquareParams.SquareSize))
	if err != nil {
		return nil, err
	}

	gameID, err := insertGame(ctx, dbConnect, createSquareParams.TeamA, createSquareParams.TeamB)
	if err != nil {
		return nil, err
	}

	err = generateFootballSquareGame(ctx, int(createSquareParams.SquareSize), dbConnect, gameID, squareID)
	if err != nil {
		return nil, err
	}

	createSquareResponse.SquareGUID = squareGuid.String()
	return &createSquareResponse, err
}

func generateFootballSquareGame(ctx context.Context, squareSize int, dbConnect *db.MySQL, gameID int64, squareID int64) error {
	for row := 1; row <= squareSize; row++ {
		for column := 1; column <= squareSize; column++ {
			_, err := dbConnect.QUERIES.CreateFootballSquareGame(ctx, db.CreateFootballSquareGameParams{
				GameID:      sql.NullInt32{Int32: int32(gameID), Valid: true},
				SquareID:    sql.NullInt32{Int32: int32(squareID), Valid: true},
				RowIndex:    sql.NullInt32{Int32: int32(row), Valid: true},
				ColumnIndex: sql.NullInt32{Int32: int32(column), Valid: true},
			})

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func insertGame(ctx context.Context, dbConnect *db.MySQL, teamA string, teamB string) (int64, error) {
	gameGuid := uuid.New()
	createGameResult, err := dbConnect.QUERIES.CreateGames(ctx, db.CreateGamesParams{
		GameGuid: gameGuid.String(),
		Sport:    sql.NullString{String: "football", Valid: true},
		TeamA:    sql.NullString{String: teamA, Valid: true},
		TeamB:    sql.NullString{String: teamB, Valid: true},
	})
	if err != nil {
		return 0, err
	}

	gameID, err := createGameResult.LastInsertId()
	if err != nil {
		return 0, err
	}
	return gameID, nil
}

func insertSquare(ctx context.Context, dbConnect *db.MySQL, squareGuid uuid.UUID, squareSize int) (int64, error) {
	createSquareResult, err := dbConnect.QUERIES.CreateSquare(ctx, db.CreateSquareParams{
		SquareGuid: squareGuid.String(),
		SquareSize: sql.NullInt32{Int32: int32(squareSize), Valid: true},
	})
	if err != nil {
		return 0, err
	}

	squareID, err := createSquareResult.LastInsertId()
	if err != nil {
		return 0, err
	}

	return squareID, nil
}
