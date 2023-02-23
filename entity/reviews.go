package entity

import "time"

type Review struct {
	Id        int
	UserId    int
	TransId   int
	Review    string
	Rating    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type InputReview struct {
	Review string `json:"review"`
	Rating int    `json:"rating" binding:"required,min=0,max=5"`
}
