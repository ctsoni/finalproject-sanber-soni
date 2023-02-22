package helper

import (
	"finalproject-sanber-soni/entity"
	"github.com/go-playground/validator/v10"
)

func FormatError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

type Response struct {
	Message string
	Code    int
	Status  string
	Data    interface{}
}

func APIResponse(message string, code int, status string, data interface{}) Response {

	response := Response{
		Message: message,
		Code:    code,
		Status:  status,
		Data:    data,
	}

	return response
}

type UserResponseFormat struct {
	ID       int
	FullName string
	Email    string
	Token    string
}

func FormatUserResponse(user entity.Users, token string) UserResponseFormat {
	response := UserResponseFormat{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Token:    token,
	}

	return response
}

type UserEditResponseFormat struct {
	ID       int
	FullName string
	Email    string
}

func FormatUserEditResponse(user entity.Users) UserEditResponseFormat {
	response := UserEditResponseFormat{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
	}

	return response
}
