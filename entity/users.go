package entity

import "time"

type Users struct {
	ID           int       `json:"id"`
	FullName     string    `json:"full_name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	IsAdmin      bool      `json:"is_admin"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type InputRegisterUsers struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type InputLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
