package routes

import (
	"market/web/handlers"
	"market/web/handlers/middlewares"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, userHandler *handlers.UserHandler, authHandler *handlers.AuthHandler, itemHandler *handlers.ItemHandler, dealHandler *handlers.DealHandler) {
	e.POST("/login", authHandler.Login)
	e.POST("/register", userHandler.CreateUser)
	e.POST("/refresh", authHandler.RefreshToken)

	authGroup := e.Group("/auth")
	authGroup.Use(middlewares.JWTMiddleware)

	InitUserRoutes(authGroup, userHandler)
	InitItemRoutes(authGroup, itemHandler)
	InitDealRoutes(authGroup, dealHandler)
}

func InitUserRoutes(group *echo.Group, handler *handlers.UserHandler) {
	group.GET("/users/:id", handler.GetUser)
	group.GET("/users", handler.GetUsers)
	group.PUT("/users/:id", handler.UpdateUser)
	group.DELETE("/users/:id", handler.DeleteUser)
}

func InitItemRoutes(group *echo.Group, handler *handlers.ItemHandler) {
	group.GET("/items/:id", handler.GetItem)
	group.GET("/items", handler.GetItems)
	group.POST("/items", handler.CreateItem)
	group.PUT("/items/:id", handler.UpdateItem)
	group.DELETE("/items/:id", handler.DeleteItem)
}

func InitDealRoutes(group *echo.Group, handler *handlers.DealHandler) {
	group.GET("/deals/:id", handler.GetDeal)
	group.GET("/deals", handler.GetDeals)
	group.POST("/deals", handler.CreateDeal)
	group.PUT("/deals/:id", handler.UpdateDeal)
	group.DELETE("/deals/:id", handler.DeleteDeal)
}
