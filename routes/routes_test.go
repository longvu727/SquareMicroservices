package routes

import (
	"bytes"
	"errors"
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

func TestRoutesTestSuite(t *testing.T) {
	suite.Run(t, new(RoutesTestSuite))
}
func (suite *RoutesTestSuite) getTestError() error {
	return errors.New("test error")
}

func (suite *RoutesTestSuite) TestCreateSquare() {

	url := "/CreateSquare"
	ctrl := gomock.NewController(suite.T())

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(`{"side_length":10}`)))
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	mockSquare := mocksquareapp.NewMockSquare(ctrl)
	mockSquare.EXPECT().
		CreateDBSquare(gomock.Any(), gomock.Any()).
		Times(1).
		Return(&app.CreateSquareResponse{SquareID: 10, SquareGUID: uuid.NewString()}, nil)

	routes := Routes{Apps: mockSquare}
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(http.MethodPost+" "+url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusOK)
}

func (suite *RoutesTestSuite) TestCreateSquareError() {

	url := "/CreateSquare"
	ctrl := gomock.NewController(suite.T())

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(`{}`)))
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	mockSquare := mocksquareapp.NewMockSquare(ctrl)
	mockSquare.EXPECT().
		CreateDBSquare(gomock.Any(), gomock.Any()).
		Times(1).
		Return(&app.CreateSquareResponse{}, suite.getTestError())

	routes := Routes{Apps: mockSquare}
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(http.MethodPost+" "+url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusInternalServerError)
}

func (suite *RoutesTestSuite) TestGetSquare() {

	url := "/GetSquare"
	ctrl := gomock.NewController(suite.T())

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(`{"square_id":10}`)))
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
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(http.MethodPost+" "+url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusOK)
}

func (suite *RoutesTestSuite) TestGetSquareError() {

	url := "/GetSquare"
	ctrl := gomock.NewController(suite.T())

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(`{}`)))
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	mockSquare := mocksquareapp.NewMockSquare(ctrl)
	mockSquare.EXPECT().
		GetDBSquare(gomock.Any(), gomock.Any()).
		Times(1).
		Return(&app.GetSquareResponse{}, suite.getTestError())

	routes := Routes{Apps: mockSquare}
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(http.MethodPost+" "+url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusInternalServerError)
}

func (suite *RoutesTestSuite) TestHome() {

	url := "/"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	routes := NewRoutes()
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusOK)
}
