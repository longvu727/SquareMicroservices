package db

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SquareTestSuite struct {
	suite.Suite
}

func (suite *SquareTestSuite) SetupTest() {}

func TestSquareTestSuite(t *testing.T) {
	suite.Run(t, new(SquareTestSuite))
}
