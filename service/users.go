package service

import (
	"errors"
	"finalproject-sanber-soni/entity"
	"finalproject-sanber-soni/repository"
	"golang.org/x/crypto/bcrypt"
)

// UserService is contract to interact with UserRepository
type UserService interface {
	RegisterUser(input entity.InputRegisterUsers) (entity.Users, error)
	Login(input entity.InputLogin) (entity.Users, error)
}

// userService is object that has userRepository field with type repository.UserRepository interface contract
type userService struct {
	userRepository repository.UserRepository
}

// NewUserService create new userService obejct with userRepository value
func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

// RegisterUser is userService method to map input from user into entity.Users object
// and pass it into userRepository.Save()
func (s *userService) RegisterUser(input entity.InputRegisterUsers) (entity.Users, error) {
	var user entity.Users

	user.FullName = input.FullName
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)

	// checking if email already exist
	_, emailExist, err := s.userRepository.FindByEmail(user.Email)
	if err != nil {
		return user, err
	}

	// if email not available or email already exist
	if emailExist {
		return user, errors.New("email already exist")
	}

	// if email available
	newUser, err := s.userRepository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

// Login is userService method to map input from user into entity.Users object
// and pass it to check if the email exist on database
// then compare the user input password with password hash in database
func (s *userService) Login(input entity.InputLogin) (entity.Users, error) {
	email := input.Email
	pwd := input.Password

	// check if the email exist
	user, _, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return user, errors.New("user not found")
	}

	if user.ID == 0 {
		return user, errors.New("user not found")
	}

	// compare user input password with password hash in database
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(pwd))
	if err != nil {
		return user, err
	}

	return user, nil
}
