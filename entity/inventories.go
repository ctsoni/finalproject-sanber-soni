package entity

import "time"

type Inventory struct {
	Id          int
	CatId       int
	Name        string
	Description string
	IsAvailable bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type InputInventory struct {
	Id           int
	CatId        int    `json:"cat_id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	IsAvailable  bool   `json:"is_available"`
	StockUnit    int    `json:"stock_unit"`
	PricePerUnit int    `json:"price_per_unit"`
}

type UpdateInventory struct {
	Id           int
	CatId        int    `json:"cat_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	IsAvailable  bool   `json:"is_available"`
	StockUnit    int    `json:"stock_unit"`
	PricePerUnit int    `json:"price_per_unit"`
}

type Stock struct {
	Id           int
	InvenId      int
	StockUnit    int
	PricePerUnit int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type InventoryStock struct {
	InventoryId    int
	CatId          int
	Name           string
	Description    string
	IsAvailable    bool
	InvenCreatedAt time.Time
	InvenUpdatedAt time.Time
	StockId        int
	InvenId        int
	StockUnit      int
	PricePerUnit   int
	StockCreatedAt time.Time
	StockUpdatedAt time.Time
}
