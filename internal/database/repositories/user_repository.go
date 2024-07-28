package repositories

import (
	"market/internal/database"
	"market/internal/database/models"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	Create(user models.NewUser) (models.User, error)
	Get(id int) (models.User, error)
	GetAll(page database.PageInfo) ([]models.User, error)
	GetByUsername(username string) (models.User, error)
	Update(user models.User) error
	Delete(id int) error
}

type UserRepository struct {
	DB *sqlx.DB
}

func (repo *UserRepository) Create(newUser models.NewUser) (models.User, error) {
	query := "INSERT INTO users (username, email, password, salt) VALUES ($1, $2, $3, $4) returning id"

	var userId int
	err := repo.DB.QueryRow(query, newUser.Username, newUser.Email, newUser.Password, newUser.Salt).Scan(&userId)

	return models.User{Id: userId, Username: newUser.Username, Email: newUser.Email, Password: newUser.Password, Salt: newUser.Salt}, err
}

func (repo *UserRepository) Get(id int) (models.User, error) {
	query := "SELECT * FROM users WHERE id = $1"

	var user models.User
	err := repo.DB.Get(&user, query, id)

	return user, err
}

func (repo *UserRepository) GetAll(page database.PageInfo) ([]models.User, error) {
	query := "SELECT * FROM users LIMIT $1 OFFSET $2"

	offset := page.Offset()

	var users []models.User
	err := repo.DB.Select(&users, query, page.PageSize, offset)

	return users, err
}

func (repo *UserRepository) Update(user models.User) error {
	query := "UPDATE users SET username = $1, email = $2, password = $3, salt = $4 WHERE id = $5"

	_, err := repo.DB.Exec(query, user.Username, user.Email, user.Password, user.Salt, user.Id)

	return err
}

func (repo *UserRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id = $1"

	_, err := repo.DB.Exec(query, id)

	return err
}

func (repo *UserRepository) GetByUsername(username string) (models.User, error) {
	query := "SELECT * FROM users WHERE username = $1"

	var user models.User
	err := repo.DB.Get(&user, query, username)

	return user, err
}
