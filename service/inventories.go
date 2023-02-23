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

	_, found, _ := s.repository.FindByName(input.Name)
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
	if !isAdmin {
		return entity.Inventory{}, entity.Stock{}, errors.New("you're not authorized")
	}

	realInput, idExsit, _ := s.repository.FindById(id)
	if !idExsit {
		return entity.Inventory{}, entity.Stock{}, errors.New("inventory id not found")
	}

	if input.CatId != realInput.CatId {
		realInput.CatId = input.CatId
	}

	if input.Name != "" {
		_, found, _ := s.repository.FindByName(input.Name)
		if found {
			return entity.Inventory{}, entity.Stock{}, errors.New("name already exist")
		}
		realInput.Name = input.Name
	}

	if input.Description != "" {
		realInput.Description = input.Description
	}

	if realInput.IsAvailable != input.IsAvailable {
		realInput.IsAvailable = input.IsAvailable
	}

	if input.StockUnit != realInput.StockUnit {
		realInput.StockUnit = input.StockUnit
	}

	if input.PricePerUnit != realInput.PricePerUnit {
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

	_, idExist, _ := s.repository.FindById(id)
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
	_, idExist, _ := s.repository.FindById(id)
	if !idExist {
		return entity.InputInventory{}, errors.New("inventory id not found")
	}

	inventory, err := s.repository.GetById(id)
	if err != nil {
		return inventory, err
	}

	return inventory, nil
}
