package repository

import (
	"database/sql"
	"finalproject-sanber-soni/entity"
	"time"
)

type ReviewRepository interface {
	Save(review entity.Review) (entity.Review, error)
	FindById(id int) (entity.Review, error)
	FindTransByTransId(transID int) (entity.Transaction, error)
	Update(review entity.Review) (entity.Review, error)
	Delete(review entity.Review) error
	GetAll() ([]entity.Review, error)
	GetByUserId(userID int) ([]entity.Review, error)
	GetByInvenId(invenID int) ([]entity.Review, error)
	FindItemByTransId(transID int) (string, error)
}

type reviewRepository struct {
	db *sql.DB
}

func NewReviewRepository(db *sql.DB) *reviewRepository {
	return &reviewRepository{db}
}

func (r *reviewRepository) Save(review entity.Review) (entity.Review, error) {
	sqlStatement := `
	INSERT INTO reviews (user_id, trans_id, review, rating) 
	VALUES ($1, $2, $3, $4)
	RETURNING *`
	err := r.db.QueryRow(
		sqlStatement,
		review.UserId,
		review.TransId,
		review.Review,
		review.Rating).Scan(
		&review.Id,
		&review.UserId,
		&review.TransId,
		&review.Review,
		&review.Rating,
		&review.CreatedAt,
		&review.UpdatedAt)
	if err != nil {
		return review, err
	}

	return review, nil
}

func (r *reviewRepository) FindById(id int) (entity.Review, error) {
	var review entity.Review

	sqlStatement := `SELECT * FROM reviews WHERE id = $1`
	err := r.db.QueryRow(
		sqlStatement,
		id).Scan(
		&review.Id,
		&review.UserId,
		&review.TransId,
		&review.Review,
		&review.Rating,
		&review.CreatedAt,
		&review.UpdatedAt)

	if err != nil {
		return review, err
	}

	return review, nil
}

func (r *reviewRepository) Update(review entity.Review) (entity.Review, error) {
	sqlStatement := `
	UPDATE reviews
	SET review = $1, rating = $2, updated_at = $3
	WHERE id = $4
	RETURNING *`
	err := r.db.QueryRow(
		sqlStatement,
		review.Review,
		review.Rating,
		time.Now(),
		review.Id).Scan(
		&review.Id,
		&review.UserId,
		&review.TransId,
		&review.Review,
		&review.Rating,
		&review.CreatedAt,
		&review.UpdatedAt)

	if err != nil {
		return review, err
	}

	return review, nil
}

func (r *reviewRepository) Delete(review entity.Review) error {
	sqlStatement := `DELETE FROM reviews WHERE id = $1`
	err := r.db.QueryRow(sqlStatement, review.Id).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *reviewRepository) FindTransByTransId(transID int) (entity.Transaction, error) {
	var transaction entity.Transaction

	sqlStatement := `SELECT * FROM transactions WHERE id = $1`
	err := r.db.QueryRow(
		sqlStatement,
		transID).Scan(
		&transaction.Id,
		&transaction.UserId,
		&transaction.InvenId,
		&transaction.Unit,
		&transaction.TotalPrice,
		&transaction.Status,
		&transaction.StartAt,
		&transaction.FinishAt,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
		&transaction.ExpiredAt,
		&transaction.StockRetrieved)

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *reviewRepository) GetAll() ([]entity.Review, error) {
	var result []entity.Review

	sqlStatement := "SELECT * FROM reviews"
	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		return result, err
	}

	defer rows.Close()
	for rows.Next() {
		var review entity.Review
		err = rows.Scan(&review.Id,
			&review.UserId,
			&review.TransId,
			&review.Review,
			&review.Rating,
			&review.CreatedAt,
			&review.UpdatedAt)
		if err != nil {
			return result, err
		}

		result = append(result, review)
	}

	return result, nil
}

func (r *reviewRepository) GetByUserId(userID int) ([]entity.Review, error) {
	var result []entity.Review

	sqlStatement := "SELECT * FROM reviews WHERE user_id = $1"
	rows, err := r.db.Query(sqlStatement, userID)
	if err != nil {
		return result, err
	}

	defer rows.Close()
	for rows.Next() {
		var review entity.Review
		err = rows.Scan(
			&review.Id,
			&review.UserId,
			&review.TransId,
			&review.Review,
			&review.Rating,
			&review.CreatedAt,
			&review.UpdatedAt)
		if err != nil {
			return result, err
		}

		result = append(result, review)
	}

	return result, nil
}

func (r *reviewRepository) GetByInvenId(invenID int) ([]entity.Review, error) {
	var result []entity.Review

	sqlStatement := `
	SELECT r.id, r.user_id, r.trans_id, r.review, r.rating, r.created_at, r.updated_at FROM reviews r 
	    JOIN transactions t ON r.trans_id = t.id
	    JOIN inventories i ON t.inven_id = i.id
	    WHERE i.id = $1`
	rows, err := r.db.Query(sqlStatement, invenID)
	if err != nil {
		return result, err
	}

	defer rows.Close()
	for rows.Next() {
		var review entity.Review
		err = rows.Scan(
			&review.Id,
			&review.UserId,
			&review.TransId,
			&review.Review,
			&review.Rating,
			&review.CreatedAt,
			&review.UpdatedAt)
		if err != nil {
			return result, err
		}

		result = append(result, review)
	}

	return result, nil
}

func (r *reviewRepository) FindItemByTransId(transID int) (string, error) {
	var item string

	sqlStatement := `
	SELECT i.name FROM inventories i 
	    JOIN transactions t ON i.id = t.inven_id
	    WHERE t.id = $1`

	err := r.db.QueryRow(sqlStatement, transID).Scan(&item)
	if err != nil {
		return item, err
	}

	return item, nil
}
