package repositories

import (
	"market/internal/database"
	"market/internal/database/models"

	"github.com/jmoiron/sqlx"
)

type ItemRepo interface {
	Create(item models.NewItem) (models.Item, error)
	Get(id int) (models.Item, error)
	GetAll(page database.PageInfo) ([]models.Item, error)
	Update(item models.Item) error
	Delete(id int) error
}

type ItemRepository struct {
	DB *sqlx.DB
}

func (repo *ItemRepository) Create(newItem models.NewItem) (models.Item, error) {
	query := "INSERT INTO items (name, price, owner_id) VALUES ($1, $2, $3) returning id"

	var itemId int
	err := repo.DB.QueryRow(query, newItem.Name, newItem.Price, newItem.OwnerId).Scan(&itemId)

	return models.Item{Id: itemId, Name: newItem.Name, Price: newItem.Price, OwnerId: newItem.OwnerId}, err
}

func (repo *ItemRepository) Update(item models.Item) error {
	query := "UPDATE items SET name = $1, price = $2 WHERE id = $3"

	_, err := repo.DB.Exec(query, item.Name, item.Price, item.Id)

	return err
}

func (repo *ItemRepository) Get(id int) (models.Item, error) {
	query := "SELECT * FROM items WHERE id = $1"

	var item models.Item
	err := repo.DB.Get(&item, query, id)

	return item, err
}

func (repo *ItemRepository) GetAll(page database.PageInfo) ([]models.Item, error) {
	query := "SELECT * FROM items LIMIT $1 OFFSET $2"

	offset := page.Offset()

	var items []models.Item
	err := repo.DB.Select(&items, query, page.PageSize, offset)

	return items, err
}

func (repo *ItemRepository) Delete(id int) error {
	query := "DELETE FROM items WHERE id = $1"

	_, err := repo.DB.Exec(query, id)

	return err
}
