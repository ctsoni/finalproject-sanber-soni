package repository

import (
	"database/sql"
	"finalproject-sanber-soni/entity"
)

type CategoriesRepository interface {
	GetAll() ([]entity.Categories, error)
}

type categoriesRepository struct {
	db *sql.DB
}

func NewCategoriesRepository(db *sql.DB) *categoriesRepository {
	return &categoriesRepository{db}
}

func (r *categoriesRepository) GetAll() ([]entity.Categories, error) {
	var result []entity.Categories

	sqlStatement := `SELECT * FROM inventory_categories`
	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		return result, err
	}

	defer rows.Close()
	for rows.Next() {
		var cat entity.Categories
		err = rows.Scan(&cat.Id, &cat.Name, &cat.CreatedAt, &cat.UpdatedAt)
		if err != nil {
			return result, err
		}
		result = append(result, cat)
	}

	return result, nil
}
