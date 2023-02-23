package repository

import (
	"database/sql"
	"finalproject-sanber-soni/entity"
	"time"
)

type InventoryRepository interface {
	Save(inputInventory entity.InputInventory) (entity.Inventory, entity.Stock, error)
	FindById(id int) (bool, error)
	FindByName(name string) (bool, error)
	Update(inputInventory entity.UpdateInventory) (entity.Inventory, entity.Stock, error)
	Delete(inventory entity.Inventory) error
	GetAll() ([]entity.InputInventory, error)
	GetById(id int) (entity.InputInventory, error)
}

type inventoryRepository struct {
	db *sql.DB
}

func NewInventoryRepository(db *sql.DB) *inventoryRepository {
	return &inventoryRepository{db}
}

func (r *inventoryRepository) Save(inputInventory entity.InputInventory) (entity.Inventory, entity.Stock, error) {
	var inventory entity.Inventory
	var stock entity.Stock

	sqlStatement := `
	INSERT INTO inventories(cat_id, name, description, is_available)
	VALUES ($1, $2, $3, $4)
	RETURNING id, cat_id, name, description, is_available`

	sqlStatementStock := `
	INSERT INTO inventory_stocks(inven_id, stock_unit, price_per_unit)
	VALUES ($1, $2, $3)
	RETURNING stock_unit, price_per_unit`

	err := r.db.QueryRow(
		sqlStatement,
		inputInventory.CatId,
		inputInventory.Name,
		inputInventory.Description,
		inputInventory.IsAvailable).Scan(
		&inventory.Id,
		&inventory.CatId,
		&inventory.Name,
		&inventory.Description,
		&inventory.IsAvailable)
	if err != nil {
		return inventory, stock, err
	}

	err = r.db.QueryRow(
		sqlStatementStock,
		inventory.Id,
		inputInventory.StockUnit,
		inputInventory.PricePerUnit).Scan(
		&stock.StockUnit,
		&stock.PricePerUnit)
	if err != nil {
		return inventory, stock, err
	}

	return inventory, stock, nil
}

func (r *inventoryRepository) FindById(id int) (bool, error) {
	sqlStatement := `SELECT * FROM inventories WHERE id = $1`
	err := r.db.QueryRow(sqlStatement, id)
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (r *inventoryRepository) FindByName(name string) (bool, error) {
	sqlStatement := `SELECT * FROM inventories WHERE name = $1`
	err := r.db.QueryRow(sqlStatement, name)
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (r *inventoryRepository) Update(input entity.UpdateInventory) (entity.Inventory, entity.Stock, error) {
	var inventory entity.Inventory
	var stock entity.Stock

	sqlStatement := `
	UPDATE inventories 
	SET cat_id=$1, name=$2, description=$3, is_available=$4, updated_at=$5
	WHERE id = $6
	RETURNING id, cat_id, name, description, is_available`

	sqlStatementStock := `
	UPDATE inventory_stocks
	SET stock_unit=$1, price_per_unit=$2, updated_at=$3
	WHERE inven_id = $4
	RETURNING stock_uni, price_per_unit`

	err := r.db.QueryRow(
		sqlStatement,
		input.CatId,
		input.Name,
		input.Description,
		input.IsAvailable,
		time.Now(),
		input.Id).Scan(
		&inventory.Id,
		&inventory.CatId,
		&inventory.Name,
		&inventory.Description,
		&inventory.IsAvailable)

	if err != nil {
		return inventory, stock, err
	}

	err = r.db.QueryRow(
		sqlStatementStock,
		input.StockUnit,
		input.PricePerUnit,
		time.Now()).Scan(
		&inventory.Id,
		&input.StockUnit,
		&input.PricePerUnit)

	if err != nil {
		return inventory, stock, err
	}

	return inventory, stock, nil
}

func (r *inventoryRepository) Delete(inventory entity.Inventory) error {
	sqlStatement := `DELETE FROM inventories WHERE id = $1`

	err := r.db.QueryRow(sqlStatement, inventory.Id).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r inventoryRepository) GetAll() ([]entity.InputInventory, error) {
	var result []entity.InputInventory

	sqlStatement := `
	SELECT i.id, i.cat_id, i.name, i.description, i.is_available, s.stock_unit, s.price_per_unit
	FROM inventories i JOIN inventory_stocks s 
	    ON i.id = s.inven_id`

	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		return result, err
	}

	rows.Close()
	for rows.Next() {
		var inventory entity.InputInventory
		err = rows.Scan(
			&inventory.Id,
			&inventory.CatId,
			&inventory.Name,
			&inventory.Description,
			&inventory.IsAvailable,
			&inventory.StockUnit,
			&inventory.PricePerUnit)
		if err != nil {
			return result, err
		}

		result = append(result, inventory)
	}

	return result, nil
}

func (r inventoryRepository) GetById(id int) (entity.InputInventory, error) {
	var inventory entity.InputInventory

	sqlStatement := `
	SELECT i.id, i.cat_id, i.name, i.description, i.is_available, s.stock_unit, s.price_per_unit
	FROM inventories i JOIN inventory_stocks s 
	    ON i.id = s.inven_id
	    WHERE i.id = $1`

	err := r.db.QueryRow(
		sqlStatement,
		id).Scan(
		&inventory.Id,
		&inventory.CatId,
		&inventory.Name,
		&inventory.Description,
		&inventory.IsAvailable,
		&inventory.StockUnit,
		&inventory.PricePerUnit)

	if err != nil {
		return inventory, err
	}

	return inventory, nil
}
