package repository

import (
	"database/sql"
	"errors"
	"finalproject-sanber-soni/entity"
	"time"
)

type TransactionRepository interface {
	Save(transaction entity.Transaction) (entity.Transaction, error)
	FindInvenId(item string) (entity.Inventory, error)
	FindInvenById(invenID int) (string, error)
	GetStockAndPrice(invenId int) (entity.Stock, error)
	UpdateStock(invenId int, unit int) error
	FindById(transactionId int) (entity.Transaction, error)
	Update(transaction entity.Transaction) (entity.Transaction, error)
	GetAll(userID int) ([]entity.Transaction, error)
	GetByStatus(userID int, status string) ([]entity.Transaction, error)
	FindUserId(userID int) error
	UpdateStockRetrieved(transaction entity.Transaction) error
	GetAllAdmin() ([]entity.Transaction, error)
	GetByStatusAdmin(status string) ([]entity.Transaction, error)
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) Save(transaction entity.Transaction) (entity.Transaction, error) {
	sqlStatement := `
	INSERT INTO transactions(user_id, inven_id, unit, total_price, status, start_at, finish_at, expired_at)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING *`

	err := r.db.QueryRow(
		sqlStatement,
		transaction.UserId,
		transaction.InvenId,
		transaction.Unit,
		transaction.TotalPrice,
		transaction.Status,
		transaction.StartAt,
		transaction.FinishAt,
		transaction.ExpiredAt).Scan(
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

func (r *transactionRepository) FindInvenId(item string) (entity.Inventory, error) {
	var inventory entity.Inventory

	sqlStatement := `SELECT id, name FROM inventories WHERE name = $1`
	err := r.db.QueryRow(sqlStatement, item).Scan(&inventory.Id, &inventory.Name)
	if err != nil {
		return inventory, err
	}

	return inventory, nil
}

func (r *transactionRepository) FindInvenById(invenID int) (string, error) {
	var item string

	sqlStatement := `SELECT name FROM inventories WHERE id = $1`
	err := r.db.QueryRow(sqlStatement, invenID).Scan(&item)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (r *transactionRepository) GetStockAndPrice(invenId int) (entity.Stock, error) {
	var stock entity.Stock

	sqlStatement := `SELECT stock_unit, price_per_unit FROM inventory_stocks WHERE inven_id = $1`
	err := r.db.QueryRow(sqlStatement, invenId).Scan(&stock.StockUnit, &stock.PricePerUnit)
	if err != nil {
		return stock, err
	}

	return stock, nil
}

func (r *transactionRepository) UpdateStock(invenId int, unit int) error {
	sqlStatement := `UPDATE inventory_stocks SET stock_unit = $1, updated_at=$3 WHERE inven_id = $2`
	err := r.db.QueryRow(sqlStatement, unit, invenId, time.Now()).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *transactionRepository) FindById(transactionId int) (entity.Transaction, error) {
	var transaction entity.Transaction

	sqlStatment := `SELECT * FROM transactions WHERE id = $1`
	err := r.db.QueryRow(sqlStatment, transactionId).Scan(
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

func (r *transactionRepository) Update(transaction entity.Transaction) (entity.Transaction, error) {
	sqlStatement := `UPDATE transactions SET status = $1, updated_at=$4 WHERE id = $2 AND user_id = $3 RETURNING *`
	err := r.db.QueryRow(
		sqlStatement,
		transaction.Status,
		transaction.Id,
		transaction.UserId,
		time.Now()).Scan(
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

func (r *transactionRepository) GetAll(userID int) ([]entity.Transaction, error) {
	var result []entity.Transaction

	sqlStatement := `SELECT * FROM transactions WHERE user_id = $1`
	rows, err := r.db.Query(sqlStatement, userID)
	if err != nil {
		return result, err
	}

	defer rows.Close()
	for rows.Next() {
		var transaction entity.Transaction
		err = rows.Scan(&transaction.Id,
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
			return result, err
		}

		result = append(result, transaction)
	}

	return result, nil
}

func (r *transactionRepository) GetByStatus(userID int, status string) ([]entity.Transaction, error) {
	var result []entity.Transaction

	sqlStatement := `SELECT * FROM transactions WHERE user_id = $1 AND status = $2`
	rows, err := r.db.Query(sqlStatement, userID, status)
	if err != nil {
		return result, err
	}

	defer rows.Close()
	for rows.Next() {
		var transaction entity.Transaction
		err = rows.Scan(&transaction.Id,
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
			return result, err
		}

		result = append(result, transaction)
	}

	return result, nil
}

func (r *transactionRepository) FindUserId(userID int) error {
	sqlStatement := `SELECT user_id FROM transactions WHERE user_id = $1`
	err := r.db.QueryRow(sqlStatement, userID)
	if err != nil {
		return errors.New("user doesn't have transactions")
	}

	return nil
}

func (r *transactionRepository) UpdateStockRetrieved(transaction entity.Transaction) error {
	sqlStatemnt := `UPDATE transactions SET stock_retreived = $1, updated_at=$3 WHERE id = $2`
	err := r.db.QueryRow(sqlStatemnt, transaction.StockRetrieved, transaction.Id, time.Now()).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *transactionRepository) GetAllAdmin() ([]entity.Transaction, error) {
	var result []entity.Transaction

	sqlStatement := `SELECT * FROM transactions`
	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		return result, err
	}

	defer rows.Close()
	for rows.Next() {
		var transaction entity.Transaction
		err = rows.Scan(&transaction.Id,
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
			return result, err
		}

		result = append(result, transaction)
	}

	return result, nil
}

func (r *transactionRepository) GetByStatusAdmin(status string) ([]entity.Transaction, error) {
	var result []entity.Transaction

	sqlStatement := `SELECT * FROM transactions WHERE status = $1`
	rows, err := r.db.Query(sqlStatement, status)
	if err != nil {
		return result, err
	}

	defer rows.Close()
	for rows.Next() {
		var transaction entity.Transaction
		err = rows.Scan(&transaction.Id,
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
			return result, err
		}

		result = append(result, transaction)
	}

	return result, nil
}
