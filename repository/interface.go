package repository

import "github.com/guntoroyk/cake-store-api/entity"

// CakeRepoItf is the interface for the CakeRepo struct
type CakeRepoItf interface {
	GetCakes() ([]*entity.Cake, error)
	GetCake(id int) (*entity.Cake, error)
	CreateCake(cake *entity.Cake) (*entity.Cake, error)
	UpdateCake(cake *entity.Cake) (*entity.Cake, error)
	DeleteCake(id int) error
}
