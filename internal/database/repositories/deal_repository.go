package repositories

import (
	"market/internal/database"
	"market/internal/database/models"

	"github.com/jmoiron/sqlx"
)

type DealRepo interface {
	Create(deal models.NewDeal) (models.Deal, error)
	Get(id int) (models.Deal, error)
	GetAll(page database.PageInfo) ([]models.Deal, error)
	Update(deal models.Deal) error
	Delete(id int) error
}

type DealRepository struct {
	DB *sqlx.DB
}

func (repo *DealRepository) Create(newDeal models.NewDeal) (models.Deal, error) {
	query := "INSERT INTO deals (item_id, user_id, price) VALUES ($1, $2, $3) returning id"

	var dealId int
	err := repo.DB.QueryRow(query, newDeal.Item.Id, newDeal.User.Id, newDeal.Price).Scan(&dealId)

	return models.Deal{Id: dealId, Item: newDeal.Item, User: newDeal.User, Price: newDeal.Price}, err
}

func (repo *DealRepository) Get(id int) (models.Deal, error) {
	query := "SELECT * FROM deals WHERE id = $1"

	var deal models.Deal
	err := repo.DB.Get(&deal, query, id)

	return deal, err
}

func (repo *DealRepository) GetAll(page database.PageInfo) ([]models.Deal, error) {
	query := "SELECT * FROM deals LIMIT $1 OFFSET $2"

	offset := page.Offset()

	var deals []models.Deal
	err := repo.DB.Select(&deals, query, page.PageSize, offset)

	return deals, err
}

func (repo *DealRepository) Update(deal models.Deal) error {
	query := "UPDATE deals SET item_id = $1, user_id = $2, price = $3 WHERE id = $4"

	_, err := repo.DB.Exec(query, deal.Item.Id, deal.User.Id, deal.Price, deal.Id)

	return err
}

func (repo *DealRepository) Delete(id int) error {
	query := "DELETE FROM deals WHERE id = $1"

	_, err := repo.DB.Exec(query, id)

	return err
}
