package repository

import (
	"database/sql"
	"finalproject-sanber-soni/entity"
)

// UserRepository is contract to interact with database
type UserRepository interface {
	Save(user entity.Users) (entity.Users, error)
	FindByEmail(email string) (entity.Users, bool, error)
}

// userRepository is object that has db *sql.DB value
type userRepository struct {
	db *sql.DB
}

// NewUserRepository create new userRepository object that has db *sql.DB value
func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

// Save is userRepository method to save user into database
func (r *userRepository) Save(user entity.Users) (entity.Users, error) {
	sqlStatement := `
	INSERT INTO users (full_name, email, password_hash) 
	VALUES ($1, $2, $3)
	RETURNING id, full_name, email, is_admin`

	// query the sql statement and assign the return value into user object
	err := r.db.QueryRow(
		sqlStatement,
		user.FullName,
		user.Email,
		user.PasswordHash).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.IsAdmin)
	if err != nil {
		return user, err
	}

	return user, nil
}

// FindByEmail is userRepository method to find if email already exist on database
func (r *userRepository) FindByEmail(email string) (entity.Users, bool, error) {
	var user entity.Users

	sqlStatement := `
	SELECT id, full_name, email, password_hash, is_admin 
	FROM users 
	WHERE email = $1`

	// query the sql statement and assign the return value into user object
	err := r.db.QueryRow(sqlStatement, email).Scan(&user.ID, &user.FullName, &user.Email, &user.PasswordHash, &user.IsAdmin)
	if err != nil {
		return user, false, err
	}

	return user, true, nil
}
