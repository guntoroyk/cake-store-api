package db

import (
	"fmt"
	"time"

	"github.com/guntoroyk/cake-store-api/entity"
)

var (
	selectCakesQuery = `SELECT id, title, description, rating, image, created_at, updated_at FROM cakes`
	selectCakeQuery  = `SELECT id, title, description, rating, image, created_at, updated_at FROM cakes WHERE id = ?`
	updateCakeQuery  = `UPDATE cakes SET %s WHERE id = %d`
	insertCakeQuery  = `INSERT INTO cakes (title, description, rating, image, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
	deleteCakeQuery  = `DELETE FROM cakes WHERE id = ?`
)

func buildUpdateCakeQuery(c *entity.Cake) (string, []interface{}) {
	var query string
	var args []interface{}

	if c.Title != "" {
		query += "title=?,"
		args = append(args, c.Title)
	}
	if c.Description != "" {
		query += "description=?,"
		args = append(args, c.Description)
	}
	if c.Rating != 0 {
		query += "rating=?,"
		args = append(args, c.Rating)
	}
	if c.Image != "" {
		query += "image=?,"
		args = append(args, c.Image)
	}
	query += "updated_at=?,"
	args = append(args, time.Now().Format("2006-01-02 15:04:05"))

	query = fmt.Sprintf(updateCakeQuery, query[:len(query)-1], c.ID)
	return query, args
}
