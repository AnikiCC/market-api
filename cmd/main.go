package main

import (
	"market/web/routes"
	"net/http"
	"os"

	"market/internal/database/repositories"
	"market/internal/services"
	"market/web/handlers"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		panic("DB_URL environment variable not set")
	}

	db, err := sqlx.Connect("postgres", dbUrl)
	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("Failed to initialize database")
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Market API")
	})

	userRepo := &repositories.UserRepository{DB: db}
	itemRepo := &repositories.ItemRepository{DB: db}
	dealRepo := &repositories.DealRepository{DB: db}

	userService := &services.UserServiceImpl{Repo: userRepo}
	itemService := &services.ItemServiceIml{Repo: itemRepo}
	dealService := &services.DealServiceImpl{Repo: dealRepo}

	userHandler := &handlers.UserHandler{Service: userService}
	authHandler := &handlers.AuthHandler{Service: userService}
	itemHandler := &handlers.ItemHandler{Service: itemService}
	dealHandler := &handlers.DealHandler{Service: dealService}

	routes.InitRoutes(e, userHandler, authHandler, itemHandler, dealHandler)

	e.Start(":8080")
}
