package app

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/longvu727/FootballSquaresLibs/DB/db"

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

func CreateDBSquare(ctx context.Context, request *http.Request, dbConnect *db.MySQL) (*CreateSquareResponse, error) {
	var createSquareParams CreateSquareParams
	json.NewDecoder(request.Body).Decode(&createSquareParams)

	var createSquareResponse CreateSquareResponse

	squareGuid := (uuid.New()).String()

	squareID, err := insertSquare(ctx, dbConnect, squareGuid, int(createSquareParams.SquareSize))
	if err != nil {
		return &createSquareResponse, err
	}

	createSquareResponse.SquareGUID = squareGuid
	createSquareResponse.SquareID = squareID

	return &createSquareResponse, nil
}

/*
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
*/
func insertSquare(ctx context.Context, dbConnect *db.MySQL, squareGuid string, squareSize int) (int64, error) {
	createSquareResult, err := dbConnect.QUERIES.CreateSquare(ctx, db.CreateSquareParams{
		SquareGuid: squareGuid,
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
