package handlers

import (
	"log"
	"net/http"
	"strconv"

	"market/internal/database"
	"market/internal/database/models"
	"market/internal/services"
	"market/web/handlers/middlewares"

	"github.com/labstack/echo/v4"
)

type ItemHandler struct {
	Service services.ItemService
}

func (h *ItemHandler) CreateItem(c echo.Context) error {
	var newItem models.NewItem
	if err := c.Bind(&newItem); err != nil {
		log.Printf("Invalid input data: %v", err)
		return c.JSON(http.StatusBadRequest, "Invalid input data")
	}

	userId := c.Get("userId").(int)

	createdItem, err := h.Service.Create(newItem, userId)
	if err != nil {
		log.Printf("Error creating item: %v", err)
		return c.JSON(http.StatusBadRequest, "Error creating item")
	}

	return c.JSON(http.StatusCreated, createdItem)
}

func (h *ItemHandler) GetItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid item ID: %v", err)
		return c.JSON(http.StatusBadRequest, "Invalid item ID")
	}

	item, err := h.Service.Get(int(id))
	if err != nil {
		log.Printf("Item not found: %v", err)
		return echo.NewHTTPError(http.StatusNotFound, "Item not found")
	}

	return c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) GetItems(c echo.Context) error {
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

	items, err := h.Service.GetAll(page)
	if err != nil {
		log.Printf("Error retrieving items: %v", err)
		return c.JSON(http.StatusInternalServerError, "Error retrieving items")
	}

	return c.JSON(http.StatusOK, items)
}

func (h *ItemHandler) UpdateItem(c echo.Context) error {
	var item models.Item
	if err := c.Bind(&item); err != nil {
		log.Printf("Invalid input data: %v", err)
		return c.JSON(http.StatusBadRequest, "Invalid input data")
	}

	claims, ok := c.Get("userClaims").(*middlewares.Claims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, "You don't have rights")
	}

	updatedItem, err := h.Service.Update(item, claims)
	if err != nil {
		log.Printf("Error updating item: %v", err)
		return c.JSON(http.StatusBadRequest, "Error updating item")
	}

	return c.JSON(http.StatusOK, updatedItem)
}

func (h *ItemHandler) DeleteItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid item ID: %v", err)
		return c.JSON(http.StatusBadRequest, "Invalid item ID")
	}

	claims, ok := c.Get("userClaims").(*middlewares.Claims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, "You don't have rights")
	}

	if err := h.Service.Delete(id, claims); err != nil {
		log.Printf("Item not found: %v", err)
		return c.JSON(http.StatusNotFound, "Item not found")
	}

	return c.NoContent(http.StatusOK)
}
