package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/guntoroyk/cake-store-api/entity"
	"github.com/guntoroyk/cake-store-api/repository"
)

type cakeRepo struct {
	db *sql.DB
}

// NewCakeRepo will create new an CakeRepo object representation of CakeRepoItf interface
func NewCakeRepo(db *sql.DB) repository.CakeRepoItf {
	return &cakeRepo{
		db: db,
	}
}

// GetCakes will get all cakes
func (c *cakeRepo) GetCakes() ([]*entity.Cake, error) {
	rows, err := c.db.Query(selectCakesQuery)
	if err != nil {
		log.Println("func GetCakes error, ", err.Error())
		return nil, err
	}
	defer rows.Close()

	var cakes []*entity.Cake
	for rows.Next() {
		var cake entity.Cake
		err := rows.Scan(
			&cake.ID, &cake.Title,
			&cake.Description, &cake.Rating,
			&cake.Image,
			&cake.CreatedAt, &cake.UpdatedAt,
		)
		if err != nil {
			log.Println("func GetCakes error scanning, ", err.Error())
			return nil, err
		}
		cakes = append(cakes, &cake)
	}
	return cakes, nil
}

// GetCake will get cake by id
func (c *cakeRepo) GetCake(id int) (*entity.Cake, error) {
	var cake entity.Cake
	err := c.db.QueryRow(selectCakeQuery, id).Scan(
		&cake.ID, &cake.Title,
		&cake.Description, &cake.Rating,
		&cake.Image,
		&cake.CreatedAt, &cake.UpdatedAt,
	)

	if err != nil {
		log.Println("func GetCake error, ", err.Error())
		return nil, err
	}

	return &cake, nil
}

// CreateCake will create a cake
func (c *cakeRepo) CreateCake(cake *entity.Cake) (*entity.Cake, error) {
	cake.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	cake.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	res, err := c.db.Exec(insertCakeQuery,
		cake.Title, cake.Description,
		cake.Rating, cake.Image,
		cake.CreatedAt, cake.UpdatedAt,
	)
	if err != nil {
		log.Println("func CreateCake error, ", err.Error())
		return nil, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Println("func CreateCake error, ", err.Error())
		return nil, err
	}
	cake.ID = int(lastId)
	return cake, nil
}

// UpdateCake will update a cake
func (c *cakeRepo) UpdateCake(cake *entity.Cake) (*entity.Cake, error) {
	updateQuery, fields := buildUpdateCakeQuery(cake)
	resp, err := c.db.Exec(updateQuery, fields...)
	if err != nil {
		log.Println("func UpdateCake error exec, ", err.Error())
		return nil, err
	}
	rowsAffected, err := resp.RowsAffected()
	if err != nil {
		log.Println("func UpdateCake error getting rows affected, ", err.Error())
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	return cake, nil
}

// DeleteCake will delete a cake
func (c *cakeRepo) DeleteCake(id int) error {
	resp, err := c.db.Exec(deleteCakeQuery, id)
	if err != nil {
		log.Println("func DeleteCake error exec, ", err.Error())
		return err
	}
	rowsAffected, err := resp.RowsAffected()
	if err != nil {
		log.Println("func DeleteCake error getting rows affected, ", err.Error())
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
