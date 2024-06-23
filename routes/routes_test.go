package routes

import (
	"bytes"
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

func (suite *RoutesTestSuite) TestCreateSquare() {
	bytesObj := []byte(`{"side_length":10}`)
	body := bytes.NewBuffer(bytesObj)

	req, err := http.NewRequest(http.MethodPost, "/CreateSquare", body)
	ctx := req.Context()

	if err != nil {
		suite.Fail(err.Error())
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		createSquare(w, r, nil, ctx)
	})
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		suite.Fail(fmt.Sprintf("handler returned wrong status code: got %d want %v", status, http.StatusOK))
	}

	// Check the response body is what we expect.
	expected := `CreateSquare Service Acknowledged`
	if rr.Body.String() != expected {
		suite.Fail(fmt.Sprintf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected))
	}
}

func TestRoutesTestSuite(t *testing.T) {
	suite.Run(t, new(RoutesTestSuite))
}
