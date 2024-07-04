package routes

import (
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/longvu727/FootballSquaresLibs/DB/db"
	mockdb "github.com/longvu727/FootballSquaresLibs/DB/db/mock"
	"github.com/longvu727/FootballSquaresLibs/util"
	"github.com/longvu727/FootballSquaresLibs/util/resources"
	"github.com/stretchr/testify/suite"
)

type RoutesTestSuite struct {
	suite.Suite
}

func (suite *RoutesTestSuite) SetupTest() {}

func (suite *RoutesTestSuite) TestCreateSquare() {
	mockMySQL := mockdb.NewMockMySQL(gomock.NewController(suite.T()))
	mockMySQL.EXPECT().
		CreateSquare(gomock.Any(), gomock.Any()).
		Times(1).
		Return(int64(10), nil)

	req, err := http.NewRequest(http.MethodPost, "/CreateSquare", bytes.NewBuffer([]byte(`{"side_length":10}`)))
	suite.NoError(err)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	httpRecorder := httptest.NewRecorder()

	config, err := util.LoadConfig("../env", "app", "env")
	suite.NoError(err)

	resources := resources.NewResources(config, mockMySQL, req.Context())

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		createSquare(w, r, resources)
	})
	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusOK)
}

func (suite *RoutesTestSuite) TestGetSquare() {
	mockMySQL := mockdb.NewMockMySQL(gomock.NewController(suite.T()))
	mockMySQL.EXPECT().
		GetSquare(gomock.Any(), gomock.Eq(int32(10))).
		Times(1).
		Return(db.GetSquareRow{
			SquareID:     10,
			SquareGuid:   uuid.NewString(),
			SquareSize:   sql.NullInt32{Int32: 10, Valid: true},
			RowPoints:    sql.NullString{String: "", Valid: true},
			ColumnPoints: sql.NullString{String: "", Valid: true},
		}, nil)

	req, err := http.NewRequest(http.MethodPost, "/GetSquare", bytes.NewBuffer([]byte(`{"square_id":10}`)))
	suite.NoError(err)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	httpRecorder := httptest.NewRecorder()

	config, err := util.LoadConfig("../env", "app", "env")
	suite.NoError(err)

	resources := resources.NewResources(config, mockMySQL, req.Context())

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getSquare(w, r, resources)
	})
	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusOK)
}

func TestGetSquareTestSuite(t *testing.T) {
	suite.Run(t, new(RoutesTestSuite))
}
