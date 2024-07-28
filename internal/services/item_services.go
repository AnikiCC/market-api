package services

import (
	"fmt"
	"log"
	"market/internal/database"
	"market/internal/database/models"
	"market/internal/database/repositories"
	"market/web/handlers/middlewares"
	"strings"
)

type ItemService interface {
	Create(newItem models.NewItem, userId int) (models.Item, error)
	Get(id int) (models.Item, error)
	GetAll(page database.PageInfo) ([]models.Item, error)
	Update(item models.Item, claims *middlewares.Claims) (models.Item, error)
	Delete(id int, claims *middlewares.Claims) error
}

type ItemServiceIml struct {
	Repo repositories.ItemRepo
}

func (ser *ItemServiceIml) Create(newItem models.NewItem, userId int) (models.Item, error) {
	if len(newItem.Name) == 0 {
		return models.Item{}, fmt.Errorf("invalid new item name")
	}

	if newItem.Price <= 0 {
		return models.Item{}, fmt.Errorf("item price cannot be less or equal 0")
	}

	newItem.OwnerId = userId

	newItem.Name = fixName(newItem.Name)

	createdItem, err := ser.Repo.Create(newItem)
	if err != nil {
		log.Printf("failed to create item: %v", err)
		return models.Item{}, fmt.Errorf("failed to create item")
	}

	return createdItem, nil
}

func (ser *ItemServiceIml) Get(id int) (models.Item, error) {
	if id <= 0 {
		return models.Item{}, fmt.Errorf("invalid item ID")
	}

	item, err := ser.Repo.Get(int(id))
	if err != nil {
		log.Printf("failed to get item: %v", err)
		return models.Item{}, fmt.Errorf("failed to get item")
	}

	return item, nil
}

func (ser *ItemServiceIml) GetAll(page database.PageInfo) ([]models.Item, error) {
	if page.PageNumber <= 0 || page.PageSize < 0 {
		return nil, fmt.Errorf("invalid pagination")
	}

	items, err := ser.Repo.GetAll(page)
	if err != nil {
		log.Printf("failed to get items: %v", err)
		return nil, fmt.Errorf("failed to get items")
	}

	return items, nil
}

func (ser *ItemServiceIml) Update(item models.Item, claims *middlewares.Claims) (models.Item, error) {
	if item.OwnerId != claims.UserId {
		return models.Item{}, fmt.Errorf("user does not own this item")
	}

	if len(strings.TrimSpace(item.Name)) == 0 {
		return models.Item{}, fmt.Errorf("item name cannot be empty")
	}

	if item.Price <= 0 {
		return models.Item{}, fmt.Errorf("price cannot be 0")
	}

	item.Name = fixName(item.Name)

	err := ser.Repo.Update(item)
	if err != nil {
		log.Printf("failed to update item: %v", err)
		return models.Item{}, fmt.Errorf("failed to update item")
	}

	return item, nil
}

func (ser *ItemServiceIml) Delete(id int, claims *middlewares.Claims) error {
	if id <= 0 {
		return fmt.Errorf("invalid item ID")
	}

	item, err := ser.Repo.Get(id)
	if err != nil {
		return fmt.Errorf("item not found")
	}

	if item.OwnerId != claims.UserId {
		return fmt.Errorf("you can only delete your own items")
	}

	ser.Repo.Delete(id)
	return nil
}

func fixName(itemName string) string {
	itemName = strings.ReplaceAll(itemName, "  ", " ")
	itemName = strings.ReplaceAll(itemName, "\t", "")
	itemName = strings.ReplaceAll(itemName, "\n", "")
	itemName = strings.ReplaceAll(itemName, "\r", "")
	itemName = strings.TrimSpace(itemName)
	return itemName
}
