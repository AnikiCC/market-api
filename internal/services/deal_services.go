package services

import (
	"fmt"
	"log"
	"market/internal/database"
	"market/internal/database/models"
	"market/internal/database/repositories"
	"market/web/handlers/middlewares"
)

type DealService interface {
	Create(deal models.NewDeal, userId int) (models.Deal, error)
	Get(id int) (models.Deal, error)
	GetAll(page database.PageInfo) ([]models.Deal, error)
	Update(deal models.Deal, claims *middlewares.Claims) (models.Deal, error)
	Delete(id int, claims *middlewares.Claims) error
}

type DealServiceImpl struct {
	Repo repositories.DealRepo
}

func (ser *DealServiceImpl) Create(newDeal models.NewDeal, userId int) (models.Deal, error) {
	if newDeal.Price <= 0 {
		return models.Deal{}, fmt.Errorf("price must be positive")
	}

	_, err := ser.Repo.Get(int(newDeal.Item.Id))
	if err != nil {
		log.Printf("Item not found: %v", err)
		return models.Deal{}, fmt.Errorf("item not found")
	}

	newDeal.User.Id = userId

	createdDeal, err := ser.Repo.Create(newDeal)
	if err != nil {
		log.Printf("Error creating deal: %v", err)
		return models.Deal{}, fmt.Errorf("failed to create deal")
	}

	return createdDeal, nil
}

func (ser *DealServiceImpl) Get(id int) (models.Deal, error) {
	if id <= 0 {
		return models.Deal{}, fmt.Errorf("invalid deal ID")
	}

	deal, err := ser.Repo.Get(id)
	if err != nil {
		log.Printf("Error retrieving deal: %v", err)
		return models.Deal{}, fmt.Errorf("deal not found")
	}

	return deal, nil
}

func (ser *DealServiceImpl) GetAll(page database.PageInfo) ([]models.Deal, error) {
	if page.PageNumber <= 0 || page.PageSize < 0 {
		return nil, fmt.Errorf("invalid pagination")
	}

	deals, err := ser.Repo.GetAll(page)
	if err != nil {
		log.Printf("Error retrieving deals: %v", err)
		return nil, fmt.Errorf("failed to get deals")
	}

	return deals, nil
}

func (ser *DealServiceImpl) Update(deal models.Deal, claims *middlewares.Claims) (models.Deal, error) {
	if deal.Id <= 0 {
		return models.Deal{}, fmt.Errorf("deal does not exist")
	}

	_, err := ser.Repo.Get(int(deal.Item.Id))
	if err != nil {
		log.Printf("Item not found: %v", err)
		return models.Deal{}, fmt.Errorf("item not found")
	}

	deal.User.Id = claims.UserId

	err = ser.Repo.Update(deal)
	if err != nil {
		log.Printf("Failed to update deal: %v", err)
		return models.Deal{}, fmt.Errorf("failed to update deal")
	}

	return deal, nil
}

func (ser *DealServiceImpl) Delete(id int, claims *middlewares.Claims) error {
	if id <= 0 {
		return fmt.Errorf("invalid deal ID")
	}

	deal, err := ser.Repo.Get(id)
	if err != nil {
		return fmt.Errorf("deal not found")
	}

	if deal.User.Id != claims.UserId {
		return fmt.Errorf("you can only delete your own deals")
	}

	return ser.Repo.Delete(id)
}
