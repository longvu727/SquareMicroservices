package routes

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"squaremicroservices/app"
	mocksquareapp "squaremicroservices/app/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type RoutesTestSuite struct {
	suite.Suite
}

func (suite *RoutesTestSuite) TestCreateSquare() {
	ctrl := gomock.NewController(suite.T())

	req, err := http.NewRequest(http.MethodPost, "/CreateSquare", bytes.NewBuffer([]byte(`{"side_length":10}`)))
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	mockSquare := mocksquareapp.NewMockSquare(ctrl)
	mockSquare.EXPECT().
		CreateDBSquare(gomock.Any(), gomock.Any()).
		Times(1).
		Return(&app.CreateSquareResponse{SquareID: 10, SquareGUID: uuid.NewString()}, nil)

	routes := Routes{Apps: mockSquare}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		routes.createSquare(w, r, nil)
	})
	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusOK)
}

func (suite *RoutesTestSuite) TestGetSquare() {
	ctrl := gomock.NewController(suite.T())

	req, err := http.NewRequest(http.MethodPost, "/GetSquare", bytes.NewBuffer([]byte(`{"square_id":10}`)))
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	returnSquare := &app.GetSquareResponse{}
	returnSquare.SquareID = 10
	returnSquare.SquareGUID = uuid.NewString()
	returnSquare.SquareSize = 10
	returnSquare.RowPoints = ""
	returnSquare.ColumnPoints = ""

	mockSquare := mocksquareapp.NewMockSquare(ctrl)
	mockSquare.EXPECT().
		GetDBSquare(gomock.Any(), gomock.Any()).
		Times(1).
		Return(returnSquare, nil)

	routes := Routes{Apps: mockSquare}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		routes.getSquare(w, r, nil)
	})
	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusOK)
}

func TestGetSquareTestSuite(t *testing.T) {
	suite.Run(t, new(RoutesTestSuite))
}
