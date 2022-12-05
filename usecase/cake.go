package usecase

import (
	"database/sql"

	"github.com/guntoroyk/cake-store-api/entity"
	"github.com/guntoroyk/cake-store-api/repository"
)

type cakeUsecase struct {
	cakeRepo repository.CakeRepoItf
}

// NewCakeUsecase will create new an CakeUsecase object representation of CakeUsecaseItf interface
func NewCakeUsecase(cakeRepo repository.CakeRepoItf) CakeUsecaseItf {
	return &cakeUsecase{
		cakeRepo: cakeRepo,
	}
}

// GetCakes will get all cakes
func (c *cakeUsecase) GetCakes() ([]*entity.Cake, error) {
	return c.cakeRepo.GetCakes()
}

// GetCake will get cake by id
func (c *cakeUsecase) GetCake(id int) (*entity.Cake, error) {
	cake, err := c.cakeRepo.GetCake(id)
	if err != nil && err == sql.ErrNoRows {
		return nil, entity.ErrCakeNotFound
	}
	return cake, err
}

// CreateCake will create a cake
func (c *cakeUsecase) CreateCake(cake *entity.Cake) (*entity.Cake, error) {
	return c.cakeRepo.CreateCake(cake)
}

// UpdateCake will update a cake
func (c *cakeUsecase) UpdateCake(cake *entity.Cake) (*entity.Cake, error) {
	cake, err := c.cakeRepo.UpdateCake(cake)
	if err != nil && err == sql.ErrNoRows {
		return nil, entity.ErrCakeNotFound
	}
	return cake, err
}

// DeleteCake will delete a cake
func (c *cakeUsecase) DeleteCake(id int) error {
	err := c.cakeRepo.DeleteCake(id)
	if err != nil && err == sql.ErrNoRows {
		return entity.ErrCakeNotFound
	}
	return err
}
