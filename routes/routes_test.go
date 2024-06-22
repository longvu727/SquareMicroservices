package routes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RoutesTestSuite struct {
	suite.Suite
}

func (suite *RoutesTestSuite) SetupTest() {}

func (suite *RoutesTestSuite) TestGetSquare() {
	req, err := http.NewRequest("GET", "/GetSquare", nil)
	if err != nil {
		suite.Fail(err.Error())
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(getSquare)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		suite.Fail(fmt.Sprintf("handler returned wrong status code: got %d want %v", status, http.StatusOK))
	}

	// Check the response body is what we expect.
	expected := `GetSquare Service Acknowledged`
	if rr.Body.String() != expected {
		suite.Fail(fmt.Sprintf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected))
	}
}

func TestRoutesTestSuite(t *testing.T) {
	suite.Run(t, new(RoutesTestSuite))
}
