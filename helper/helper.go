package helper

import (
	"finalproject-sanber-soni/entity"
	"github.com/go-playground/validator/v10"
	"time"
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

type UserGetAllFormat struct {
	ID                   int
	FullName, Email      string
	CreatedAt, UpdatedAt time.Time
}

func FormatUserGetAllResponse(users []entity.Users) []UserGetAllFormat {
	var response []UserGetAllFormat
	for _, user := range users {
		res := UserGetAllFormat{
			ID:        user.ID,
			FullName:  user.FullName,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		response = append(response, res)
	}

	return response
}

type CatGetAllFormat struct {
	ID   int
	Name string
}

func FormatCatGetAllResponse(cats []entity.Category) []CatGetAllFormat {
	var response []CatGetAllFormat
	for _, cat := range cats {
		res := CatGetAllFormat{
			ID:   cat.Id,
			Name: cat.Name,
		}
		response = append(response, res)
	}

	return response
}

type CatInvenGetFormat struct {
	ID, CatID         int
	Name, Description string
	IsAvailable       bool
}

func FormatCatInvenGetResponse(invens []entity.Inventory) []CatInvenGetFormat {
	var response []CatInvenGetFormat
	for _, inven := range invens {
		res := CatInvenGetFormat{
			ID:          inven.Id,
			CatID:       inven.CatId,
			Name:        inven.Name,
			Description: inven.Description,
			IsAvailable: inven.IsAvailable,
		}
		response = append(response, res)
	}

	return response
}

type InventoryFormat struct {
	ID, CatID               int
	Name, Description       string
	IsAvailable             bool
	StockUnit, PricePerUnit int
}

func FormatInventorySaveResponse(inventory entity.Inventory, stock entity.Stock) InventoryFormat {
	response := InventoryFormat{
		ID:           inventory.Id,
		CatID:        inventory.CatId,
		Name:         inventory.Name,
		Description:  inventory.Description,
		IsAvailable:  inventory.IsAvailable,
		StockUnit:    stock.StockUnit,
		PricePerUnit: stock.PricePerUnit,
	}

	return response
}

type OutputTransaction struct {
	Information string
	Id          int
	UserId      int
	Item        string
	Unit        int
	TotalPrice  int
	Status      string
	StartAt     time.Time
	FinishAt    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ExpiredAt   time.Time
}

func FormatTransactionResponse(transaction entity.Transaction, item string, info string) OutputTransaction {
	response := OutputTransaction{
		Information: info,
		Id:          transaction.Id,
		UserId:      transaction.UserId,
		Item:        item,
		Unit:        transaction.Unit,
		TotalPrice:  transaction.TotalPrice,
		Status:      transaction.Status,
		StartAt:     transaction.StartAt,
		FinishAt:    transaction.FinishAt,
		CreatedAt:   transaction.CreatedAt,
		UpdatedAt:   transaction.UpdatedAt,
		ExpiredAt:   transaction.ExpiredAt,
	}

	return response
}

type OutputReview struct {
	ID        int
	User      string
	Item      string
	TransID   int
	Review    string
	Rating    int
	UpdatedAt time.Time
}

func FormatReviewResponse(review entity.Review, user string, item string) OutputReview {
	response := OutputReview{
		ID:        review.Id,
		User:      user,
		Item:      item,
		TransID:   review.TransId,
		Review:    review.Review,
		Rating:    review.Rating,
		UpdatedAt: review.UpdatedAt,
	}

	return response
}
