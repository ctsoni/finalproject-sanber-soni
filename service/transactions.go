package service

import (
	"errors"
	"finalproject-sanber-soni/entity"
	"finalproject-sanber-soni/repository"
	"strings"
	"time"
)

type TransactionService interface {
	CreateTransaction(input entity.InputTransaction, userID int, isAdmin bool) (entity.Transaction, error)
	UpdateTransaction(id int, status string, userID int, isAdmin bool) (entity.Transaction, string, error)
	GetAll(userID int) ([]entity.Transaction, error)
	GetByStatus(userID int, status string) ([]entity.Transaction, error)
	ValidateStatus(status string) error
	TransactionDoneByAdmin(transId int, isAdmin bool) error
	GetAllAdmin(isAdmin bool) ([]entity.Transaction, error)
	GetByStatusAdmin(status string, isAdmin bool) ([]entity.Transaction, error)
}

type transactionService struct {
	repository repository.TransactionRepository
}

func NewTransactionService(repository repository.TransactionRepository) *transactionService {
	return &transactionService{repository}
}

func (s *transactionService) CreateTransaction(input entity.InputTransaction, userID int, isAdmin bool) (entity.Transaction, error) {
	var transaction entity.Transaction
	if isAdmin {
		return transaction, errors.New("admin not allowed")
	}

	item, err := s.repository.FindInvenId(input.Item)
	if err != nil {
		return transaction, errors.New("no item with that name")
	}

	stock, err := s.repository.GetStockAndPrice(item.Id)
	if err != nil {
		return transaction, err
	}

	if stock.StockUnit == 0 || stock.PricePerUnit == 0 || stock.StockUnit < input.Unit {
		return transaction, errors.New("the item currently fully booked/not enough stock unit for your booking")
	}

	if input.StartAt == input.FinishAt || input.StartAt.After(input.FinishAt) {
		return transaction, errors.New("you cannot put start time same with finish time or start time after finish time")
	}

	// TODO add stock validation by time
	transaction.UserId = userID
	transaction.InvenId = item.Id
	transaction.Unit = input.Unit
	transaction.TotalPrice = input.Unit * stock.PricePerUnit
	transaction.Status = "Unpaid"
	transaction.StartAt = input.StartAt
	transaction.FinishAt = input.FinishAt
	transaction.ExpiredAt = time.Now().Add(5 * time.Minute)

	createdTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return createdTransaction, err
	}

	currentStock := stock.StockUnit - input.Unit
	err = s.repository.UpdateStock(item.Id, currentStock)
	if err != nil {
		return createdTransaction, err
	}

	return createdTransaction, nil
}

func (s *transactionService) UpdateTransaction(id int, status string, userID int, isAdmin bool) (entity.Transaction, string, error) {
	// validate if the user admin
	if isAdmin {
		return entity.Transaction{}, "", errors.New("admin not allowed")
	}

	transaction, err := s.repository.FindById(id)
	if err != nil {
		return transaction, "", errors.New("transaction not found")
	}

	// validate if the user is the owner of transaction
	if transaction.UserId != userID {
		return entity.Transaction{}, "", errors.New("transaction not yours")
	}

	// validate if transaction already closed
	if transaction.Status != "Unpaid" {
		return entity.Transaction{}, "", errors.New("transaction already closed")
	}

	// validate if status parameter is pay or cancel
	if strings.ToLower(status) != "pay" && strings.ToLower(status) != "cancel" {
		return transaction, "", errors.New("status input not allowed")
	}

	var statusMsg string
	if strings.ToLower(status) == "pay" {
		statusMsg = "Paid"
	} else if strings.ToLower(status) == "cancel" {
		statusMsg = "Canceled"
	}

	item, err := s.repository.FindInvenById(transaction.InvenId)
	if err != nil {
		return transaction, item, err
	}

	// validate if transaction already expired
	if time.Now().After(transaction.ExpiredAt) {
		statusMsg = "Failed"
		stock, err := s.repository.GetStockAndPrice(transaction.InvenId)
		if err != nil {
			return transaction, item, err
		}
		err = s.repository.UpdateStock(transaction.InvenId, transaction.Unit+stock.StockUnit)
		if err != nil {
			return transaction, item, err
		}

		transaction.Status = statusMsg
		transaction.StockRetrieved = true
		updatedTransaction, err := s.repository.Update(transaction)
		if err != nil {
			return updatedTransaction, "", err
		}

		return updatedTransaction, item, errors.New("your transaction already expired")
	}

	// Action for canceled transaction
	if statusMsg == "Canceled" {
		stock, err := s.repository.GetStockAndPrice(transaction.InvenId)
		if err != nil {
			return transaction, item, err
		}
		err = s.repository.UpdateStock(transaction.InvenId, transaction.Unit+stock.StockUnit)
		if err != nil {
			return transaction, item, err
		}
		transaction.StockRetrieved = true
	}

	transaction.Status = statusMsg
	updatedTransaction, err := s.repository.Update(transaction)
	if err != nil {
		return updatedTransaction, "", err
	}

	return updatedTransaction, item, nil
}

func (s *transactionService) GetAll(userID int) ([]entity.Transaction, error) {
	//err := s.repository.FindUserId(userID)
	//if err != nil {
	//	return nil, err
	//}

	transactions, err := s.repository.GetAll(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *transactionService) GetByStatus(userID int, status string) ([]entity.Transaction, error) {
	//err := s.repository.FindUserId(userID)
	//if err != nil {
	//	return nil, err
	//}

	transactions, err := s.repository.GetByStatus(userID, status)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *transactionService) ValidateStatus(status string) error {
	if status != "Failed" && status != "Unpaid" && status != "Paid" && status != "Canceled" {
		return errors.New("status invalid")
	}

	return nil
}

func (s *transactionService) TransactionDoneByAdmin(transId int, isAdmin bool) error {
	if !isAdmin {
		return errors.New("you're not authorized")
	}

	transaction, err := s.repository.FindById(transId)
	if err != nil {
		return errors.New("transaction id not found")
	}

	if transaction.Status != "Paid" {
		return errors.New("transaction is not paid and stock already back up")
	}

	if transaction.StockRetrieved {
		return errors.New("the stock from this transactions already retrieved")
	}

	stock, err := s.repository.GetStockAndPrice(transaction.InvenId)
	if err != nil {
		return err
	}
	err = s.repository.UpdateStock(transaction.InvenId, transaction.Unit+stock.StockUnit)
	if err != nil {
		return err
	}

	transaction.StockRetrieved = true
	err = s.repository.UpdateStockRetrieved(transaction)
	if err != nil {
		return err
	}

	return nil
}

func (s *transactionService) GetAllAdmin(isAdmin bool) ([]entity.Transaction, error) {
	if !isAdmin {
		return nil, errors.New("you're not authorized")
	}

	transactions, err := s.repository.GetAllAdmin()
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *transactionService) GetByStatusAdmin(status string, isAdmin bool) ([]entity.Transaction, error) {
	if !isAdmin {
		return nil, errors.New("you're not authorized")
	}

	transactions, err := s.repository.GetByStatusAdmin(status)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
