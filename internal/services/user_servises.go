package services

import (
	"fmt"
	"market/internal/database"
	"market/internal/database/models"
	"market/internal/database/repositories"
	"market/web/handlers/middlewares"
	"strings"
)

type UserService interface {
	Create(newUser models.NewUser) (models.UserResponse, error)
	Get(id int) (models.UserResponse, error)
	GetAll(page database.PageInfo) ([]models.UserResponse, error)
	Update(user models.User, claims *middlewares.Claims) (models.UserResponse, error)
	Delete(id int, claims *middlewares.Claims) error
	Authenticate(username, password string) (models.UserResponse, error)
}

type UserServiceImpl struct {
	Repo repositories.UserRepo
	Pass PasswordManager
}

type UserContext struct {
	UserID int
}

func (ser *UserServiceImpl) Create(newUser models.NewUser) (models.UserResponse, error) {
	newUser.Username = fixUserName(newUser.Username)
	if len(newUser.Username) == 0 {
		return models.UserResponse{}, fmt.Errorf("invalid user name")
	}

	salt, err := ser.Pass.GenerateSalt()
	if err != nil {
		return models.UserResponse{}, fmt.Errorf("failed to generate salt")
	}

	newUser.Salt = salt

	hashedPassword, err := ser.Pass.HashPassword(newUser.Password, newUser.Salt)
	if err != nil {
		return models.UserResponse{}, fmt.Errorf("failed to hash password")
	}

	newUser.Password = hashedPassword
	createdUser, err := ser.Repo.Create(newUser)
	if err != nil {
		return models.UserResponse{}, fmt.Errorf("failed to create user")
	}

	return createdUser.ToResponse(), nil
}

func (ser *UserServiceImpl) Get(id int) (models.UserResponse, error) {
	if id <= 0 {
		return models.UserResponse{}, fmt.Errorf("invalid user ID")
	}

	user, err := ser.Repo.Get(id)
	if err != nil {
		return models.UserResponse{}, fmt.Errorf("failed to get user")
	}

	return user.ToResponse(), nil
}

func (ser *UserServiceImpl) GetAll(page database.PageInfo) ([]models.UserResponse, error) {
	if page.PageNumber <= 0 || page.PageSize < 0 {
		return nil, fmt.Errorf("invalid pagination")
	}

	users, err := ser.Repo.GetAll(page)
	if err != nil {
		return nil, fmt.Errorf("failed to get users")
	}

	userResponses := make([]models.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = user.ToResponse()
	}

	return userResponses, nil
}

func (ser *UserServiceImpl) Update(user models.User, claims *middlewares.Claims) (models.UserResponse, error) {
	if user.Id != claims.UserId {
		return models.UserResponse{}, fmt.Errorf("not authorized to update this user")
	}

	if len(strings.TrimSpace(user.Username)) == 0 {
		return models.UserResponse{}, fmt.Errorf("name cannot be empty")
	}

	hashedPassword, err := ser.Pass.HashPassword(user.Password, user.Salt)
	if err != nil {
		return models.UserResponse{}, fmt.Errorf("failed to hash password")
	}

	user.Password = hashedPassword
	user.Username = fixUserName(user.Username)

	err = ser.Repo.Update(user)
	if err != nil {
		return models.UserResponse{}, fmt.Errorf("failed to update user: %w", err)
	}

	return user.ToResponse(), nil
}

func (ser *UserServiceImpl) Delete(id int, claims *middlewares.Claims) error {
	if id != claims.UserId {
		return fmt.Errorf("not authorized to delete this user")
	}

	ser.Repo.Delete(id)
	return nil
}

func (ser *UserServiceImpl) Authenticate(username, password string) (models.UserResponse, error) {
	user, err := ser.Repo.GetByUsername(username)
	if err != nil {
		return models.UserResponse{}, fmt.Errorf("user not found")
	}

	if !ser.Pass.CheckPasswordHash(password+user.Salt, user.Password) {
		return models.UserResponse{}, fmt.Errorf("invalid password")
	}

	return models.UserResponse{Id: user.Id, Username: user.Username}, nil
}

func fixUserName(username string) string {
	username = strings.ReplaceAll(username, "  ", " ")
	username = strings.ReplaceAll(username, "\t", "")
	username = strings.ReplaceAll(username, "\n", "")
	username = strings.ReplaceAll(username, "\r", "")
	username = strings.TrimSpace(username)
	return username
}
