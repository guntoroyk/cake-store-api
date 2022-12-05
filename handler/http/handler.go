package http

import "github.com/guntoroyk/cake-store-api/usecase"

type handler struct {
	cakeUsecase usecase.CakeUsecaseItf
}

// NewHandler is a constructor for handler
func NewHandler(cakeUsecase usecase.CakeUsecaseItf) *handler {
	return &handler{
		cakeUsecase: cakeUsecase,
	}
}
