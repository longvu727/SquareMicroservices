package app

import (
	"context"
	"database/sql"
	"errors"
	"math/rand"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	db "github.com/longvu727/FootballSquaresLibs/DB/db"
	mockdb "github.com/longvu727/FootballSquaresLibs/DB/db/mock"
	"github.com/longvu727/FootballSquaresLibs/util"
	"github.com/longvu727/FootballSquaresLibs/util/resources"
	"github.com/stretchr/testify/suite"
)

type GetSquareTestSuite struct {
	suite.Suite
}

func TestGetSquareTestSuite(t *testing.T) {
	suite.Run(t, new(GetSquareTestSuite))
}

func (suite *GetSquareTestSuite) TestGetSquare() {
	randomSquare := randomSquare()

	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	mockMySQL := mockdb.NewMockMySQL(ctrl)

	mockMySQL.EXPECT().
		GetSquare(gomock.Any(), gomock.Eq(randomSquare.SquareID)).
		Times(1).
		Return(randomSquare, nil)

	config, err := util.LoadConfig("../env", "app", "env")
	suite.NoError(err)

	resources := resources.NewResources(config, mockMySQL, context.Background())

	getSquareParams := GetSquareParams{SquareID: int(randomSquare.SquareID)}
	square, err := NewSquareApp().GetDBSquare(getSquareParams, resources)
	suite.NoError(err)

	suite.Equal(randomSquare.SquareID, int32(square.SquareID))
	suite.Equal(randomSquare.SquareGuid, square.SquareGUID)
	suite.Equal(randomSquare.RowPoints.String, square.RowPoints)
	suite.Equal(randomSquare.ColumnPoints.String, square.ColumnPoints)
	suite.Equal(randomSquare.SquareSize.Int32, int32(square.SquareSize))

	suite.Greater(len(square.ToJson()), 0)
}

func (suite *GetSquareTestSuite) TestGetSquareError() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	mockMySQL := mockdb.NewMockMySQL(ctrl)

	mockMySQL.EXPECT().
		GetSquare(gomock.Any(), gomock.Any()).
		Times(1).
		Return(db.GetSquareRow{}, errors.New("test error"))

	config, err := util.LoadConfig("../env", "app", "env")
	suite.NoError(err)

	resources := resources.NewResources(config, mockMySQL, context.Background())

	_, err = NewSquareApp().GetDBSquare(GetSquareParams{}, resources)
	suite.Error(err)
}

func intsToString(numbers []int) string {
	arrayStr := ""

	for _, number := range numbers {
		arrayStr += string(rune(number)) + ","
	}

	arrayStr = strings.TrimRight(arrayStr, ",")

	return arrayStr
}

func randomSquare() db.GetSquareRow {
	squareSize := rand.Int31n(10)

	rowPointsStr := intsToString(rand.Perm(9))
	columnPointsStr := intsToString(rand.Perm(9))

	return db.GetSquareRow{
		SquareID:     rand.Int31n(1000),
		SquareGuid:   uuid.NewString(),
		SquareSize:   sql.NullInt32{Int32: squareSize, Valid: true},
		RowPoints:    sql.NullString{String: rowPointsStr, Valid: true},
		ColumnPoints: sql.NullString{String: columnPointsStr, Valid: true},
	}
}
