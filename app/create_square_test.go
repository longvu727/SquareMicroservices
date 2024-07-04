package app

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/longvu727/FootballSquaresLibs/DB/db/mock"
	"github.com/longvu727/FootballSquaresLibs/util"
	"github.com/longvu727/FootballSquaresLibs/util/resources"
	"github.com/stretchr/testify/suite"
)

type CreateSquareTestSuite struct {
	suite.Suite
}

func (suite *CreateSquareTestSuite) SetupTest() {
}

func (suite *CreateSquareTestSuite) TestCreateSquare() {
	randomSquare := randomSquare()

	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	mockMySQL := mockdb.NewMockMySQL(ctrl)

	mockMySQL.EXPECT().
		CreateSquare(gomock.Any(), gomock.Any()).
		Times(1).
		Return(int64(randomSquare.SquareID), nil)

	config, err := util.LoadConfig("../env", "app", "env")
	suite.NoError(err)

	resources := resources.NewResources(config, mockMySQL, context.Background())

	createSquareParams := CreateSquareParams{
		SquareSize: randomSquare.SquareSize.Int32,
	}
	square, err := NewSquareApp().CreateDBSquare(createSquareParams, resources)
	suite.NoError(err)

	suite.Equal(randomSquare.SquareID, int32(square.SquareID))
}

func TestCreateSquareTestSuite(t *testing.T) {
	suite.Run(t, new(CreateSquareTestSuite))
}
