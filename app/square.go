package app

import "github.com/longvu727/FootballSquaresLibs/util/resources"

type Square interface {
	GetDBSquare(getSquareParams GetSquareParams, resources *resources.Resources) (*GetSquareResponse, error)
	CreateDBSquare(createSquareParams CreateSquareParams, resources *resources.Resources) (*CreateSquareResponse, error)
}

type SquareApp struct{}

func NewSquareApp() Square {
	return &SquareApp{}
}
