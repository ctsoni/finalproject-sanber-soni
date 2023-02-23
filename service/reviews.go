package service

import (
	"errors"
	"finalproject-sanber-soni/entity"
	"finalproject-sanber-soni/repository"
)

type ReviewService interface {
	AddReview(input entity.InputReview, userID int, transID int, isAdmin bool) (entity.Review, error)
	EditReview(input entity.InputReview, reviewID int, userID int) (entity.Review, error)
	DeleteReview(reviewID, userID int) error
	FindItem(invenID int) (string, error)
	GetAll() ([]entity.Review, error)
	GetByUserId(userID int) ([]entity.Review, error)
	GetByInvenId(InvenID int) ([]entity.Review, error)
}

type reviewService struct {
	repository repository.ReviewRepository
}

func NewReviewService(reviewRepository repository.ReviewRepository) *reviewService {
	return &reviewService{reviewRepository}
}

func (s *reviewService) AddReview(input entity.InputReview, userID int, transID int, isAdmin bool) (entity.Review, error) {
	var review entity.Review
	if isAdmin {
		return review, errors.New("admin not allowed")
	}

	transaction, err := s.repository.FindTransByTransId(transID)
	if err != nil {
		return review, errors.New("transaction id not found")
	}

	if transaction.UserId != userID {
		return review, errors.New("transaction id is not yours")
	}

	review.UserId = userID
	review.TransId = transID
	review.Review = input.Review
	review.Rating = input.Rating

	newReview, err := s.repository.Save(review)
	if err != nil {
		return newReview, err
	}

	return newReview, err
}

func (s *reviewService) EditReview(input entity.InputReview, reviewID int, userID int) (entity.Review, error) {
	review, err := s.repository.FindById(reviewID)
	if err != nil {
		return review, errors.New("review id not found")
	}

	if review.UserId != userID {
		return review, errors.New("you're not authorized to edit this review")
	}

	review.Review = input.Review
	review.Rating = input.Rating

	updatedReview, err := s.repository.Update(review)
	if err != nil {
		return updatedReview, err
	}

	return updatedReview, err
}

func (s *reviewService) DeleteReview(reviewID, userID int) error {
	review, err := s.repository.FindById(reviewID)
	if err != nil {
		return errors.New("review id not found")
	}

	if review.UserId != userID {
		return errors.New("you're not authorized")
	}

	err = s.repository.Delete(review)
	if err != nil {
		return err
	}

	return nil
}

func (s *reviewService) FindItem(transID int) (string, error) {
	item, err := s.repository.FindItemByTransId(transID)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (s *reviewService) GetAll() ([]entity.Review, error) {
	reviews, err := s.repository.GetAll()
	if err != nil {
		return reviews, err
	}

	return reviews, err
}

func (s *reviewService) GetByUserId(userID int) ([]entity.Review, error) {
	reviews, err := s.repository.GetByUserId(userID)
	if err != nil {
		return reviews, err
	}

	return reviews, err
}

func (s *reviewService) GetByInvenId(InvenID int) ([]entity.Review, error) {
	reviews, err := s.repository.GetByInvenId(InvenID)
	if err != nil {
		return reviews, err
	}

	return reviews, err
}
