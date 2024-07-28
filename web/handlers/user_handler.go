package handlers

import (
	"net/http"
	"strconv"

	"market/internal/database"
	"market/internal/database/models"
	"market/internal/services"
	"market/web/handlers/middlewares"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service services.UserService
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var newUser models.NewUser
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	createdUser, err := h.Service.Create(newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	user, err := h.Service.Get(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	pageNum, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || pageNum <= 0 {
		pageNum = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	page := database.PageInfo{
		PageNumber: pageNum,
		PageSize:   pageSize,
	}

	users, err := h.Service.GetAll(page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	claims, ok := c.Get("userClaims").(*middlewares.Claims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, "You don't have rights")
	}

	updatedUser, err := h.Service.Update(user, claims)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, updatedUser)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	claims, ok := c.Get("userClaims").(*middlewares.Claims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, "You don't have rights")
	}

	if err := h.Service.Delete(id, claims); err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
