package service

import (
	"finalproject-sanber-soni/entity"
	"finalproject-sanber-soni/repository"
)

type CategoriesService interface {
	GetAll() ([]entity.Categories, error)
}

type categoriesService struct {
	catRepository repository.CategoriesRepository
}

func NewCategoryService(categoriesRepository repository.CategoriesRepository) *categoriesService {
	return &categoriesService{categoriesRepository}
}

func (s *categoriesService) GetAll() ([]entity.Categories, error) {
	categories, err := s.catRepository.GetAll()
	if err != nil {
		return categories, err
	}

	return categories, nil
}
