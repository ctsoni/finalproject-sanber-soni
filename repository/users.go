package repository

import (
	"database/sql"
	"finalproject-sanber-soni/entity"
	"time"
)

type UserRepository interface {
	Save(user entity.Users) (entity.Users, error)
	FindByEmail(email string) (entity.Users, bool, error)
	FindById(id int) (entity.Users, error)
	Update(user entity.Users) (entity.Users, error)
	Delete(user entity.Users) error
	GetAll() ([]entity.Users, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

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

func (r *userRepository) FindById(id int) (entity.Users, error) {
	var user entity.Users

	sqlStatement := `SELECT id, full_name, email, password_hash, is_admin FROM users WHERE id = $1`
	err := r.db.QueryRow(
		sqlStatement,
		id).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.PasswordHash,
		&user.IsAdmin)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Update(user entity.Users) (entity.Users, error) {

	sqlStatement := `
	UPDATE users
	SET full_name=$2, email=$3, password_hash=$4, updated_at=$5
	WHERE id = $1
	RETURNING id, full_name, email`

	err := r.db.QueryRow(
		sqlStatement,
		user.ID,
		user.FullName,
		user.Email,
		user.PasswordHash,
		time.Now()).Scan(
		&user.ID,
		&user.FullName,
		&user.Email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Delete(user entity.Users) error {
	sqlStatement := `
	DELETE FROM users
	WHERE id = $1`

	err := r.db.QueryRow(sqlStatement, user.ID).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetAll() ([]entity.Users, error) {
	var result []entity.Users

	sqlStatement := `SELECT * FROM users`
	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		return result, err
	}

	defer rows.Close()
	for rows.Next() {
		var user entity.Users
		err = rows.Scan(
			&user.ID,
			&user.FullName,
			&user.Email,
			&user.PasswordHash,
			&user.IsAdmin,
			&user.CreatedAt,
			&user.UpdatedAt)
		if err != nil {
			return result, err
		}
		if !user.IsAdmin {
			result = append(result, user)
		}
	}

	return result, nil
}
