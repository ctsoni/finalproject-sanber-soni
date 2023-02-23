package entity

import "time"

type Transaction struct {
	Id             int
	UserId         int
	InvenId        int
	Unit           int
	TotalPrice     int
	Status         string
	StartAt        time.Time
	FinishAt       time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ExpiredAt      time.Time
	StockRetrieved bool
}

type InputTransaction struct {
	Item     string    `json:"item" binding:"required"`
	Unit     int       `json:"unit" binding:"required,gt=1"`
	StartAt  time.Time `json:"start_at" binding:"required"`
	FinishAt time.Time `json:"finish_at" binding:"required"`
}
