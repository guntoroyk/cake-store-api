package usecase

import "github.com/guntoroyk/cake-store-api/entity"

// CakeUsecaseItf is the interface for the CakeUsecase struct
type CakeUsecaseItf interface {
	GetCakes() ([]*entity.Cake, error)
	GetCake(id int) (*entity.Cake, error)
	CreateCake(cake *entity.Cake) (*entity.Cake, error)
	UpdateCake(cake *entity.Cake) (*entity.Cake, error)
	DeleteCake(id int) error
}
