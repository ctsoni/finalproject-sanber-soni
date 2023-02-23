package service

import (
	"errors"
	"finalproject-sanber-soni/entity"
	"finalproject-sanber-soni/repository"
)

type InventoryService interface {
	AddItem(inputInventory entity.InputInventory, isAdmin bool) (entity.Inventory, entity.Stock, error)
	UpdateItem(input entity.UpdateInventory, isAdmin bool, id int) (entity.Inventory, entity.Stock, error)
	DeleteItem(isAdmin bool, id int) error
	GetAll() ([]entity.InputInventory, error)
	GetById(id int) (entity.InputInventory, error)
}

type inventoryService struct {
	repository repository.InventoryRepository
}

func NewInventoryService(repository repository.InventoryRepository) *inventoryService {
	return &inventoryService{repository}
}

func (s *inventoryService) AddItem(input entity.InputInventory, isAdmin bool) (entity.Inventory, entity.Stock, error) {
	if !isAdmin {
		return entity.Inventory{}, entity.Stock{}, errors.New("you're not authorize")
	}

	found, _ := s.repository.FindByName(input.Name)
	if found {
		return entity.Inventory{}, entity.Stock{}, errors.New("name already exist")
	}

	inventory, stock, err := s.repository.Save(input)
	if err != nil {
		return inventory, stock, err
	}

	return inventory, stock, nil
}

func (s *inventoryService) UpdateItem(input entity.UpdateInventory, isAdmin bool, id int) (entity.Inventory, entity.Stock, error) {
	var realInput entity.UpdateInventory

	if !isAdmin {
		return entity.Inventory{}, entity.Stock{}, errors.New("you're not authorized")
	}

	idExsit, _ := s.repository.FindById(id)
	if !idExsit {
		return entity.Inventory{}, entity.Stock{}, errors.New("inventory id not found")
	}

	found, _ := s.repository.FindByName(input.Name)
	if found {
		return entity.Inventory{}, entity.Stock{}, errors.New("name already exist")
	}

	realInput.Id = id

	if input.CatId != 0 {
		realInput.CatId = input.CatId
	}

	if input.Name != "" {
		realInput.Name = input.Name
	}

	if input.Description != "" {
		realInput.Description = input.Description
	}

	realInput.IsAvailable = input.IsAvailable

	if input.StockUnit != 0 {
		realInput.StockUnit = input.StockUnit
	}

	if input.PricePerUnit != 0 {
		realInput.PricePerUnit = input.PricePerUnit
	}

	inventory, stock, err := s.repository.Update(realInput)
	if err != nil {
		return inventory, stock, err
	}

	return inventory, stock, nil
}

func (s *inventoryService) DeleteItem(isAdmin bool, id int) error {
	var inventory entity.Inventory

	if !isAdmin {
		return errors.New("you're not authorized")
	}

	idExist, _ := s.repository.FindById(id)
	if !idExist {
		return errors.New("inventory id not found")
	}

	inventory.Id = id

	err := s.repository.Delete(inventory)
	if err != nil {
		return err
	}

	return nil
}

func (s *inventoryService) GetAll() ([]entity.InputInventory, error) {
	inventories, err := s.repository.GetAll()
	if err != nil {
		return inventories, err
	}

	return inventories, nil
}

func (s *inventoryService) GetById(id int) (entity.InputInventory, error) {
	idExist, _ := s.repository.FindById(id)
	if !idExist {
		return entity.InputInventory{}, errors.New("inventory id not found")
	}

	inventory, err := s.repository.GetById(id)
	if err != nil {
		return inventory, err
	}

	return inventory, nil
}
