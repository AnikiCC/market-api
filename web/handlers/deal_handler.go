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

type DealHandler struct {
	Service services.DealService
}

func (h *DealHandler) CreateDeal(c echo.Context) error {
	var newDeal models.NewDeal
	if err := c.Bind(&newDeal); err != nil {
		log.Printf("Invalid input data: %v", err)
		return c.JSON(http.StatusBadRequest, "Invalid input data")
	}

	userId, ok := c.Get("userId").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	createdDeal, err := h.Service.Create(newDeal, userId)
	if err != nil {
		log.Printf("Error creating deal: %v", err)
		return c.JSON(http.StatusBadRequest, "Error creating deal")
	}

	return c.JSON(http.StatusCreated, createdDeal)
}

func (h *DealHandler) GetDeal(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid deal ID: %v", err)
		return c.JSON(http.StatusBadRequest, "Invalid deal ID")
	}

	deal, err := h.Service.Get(id)
	if err != nil {
		log.Printf("Deal not found: %v", err)
		return echo.NewHTTPError(http.StatusNotFound, "Deal not found")
	}

	return c.JSON(http.StatusOK, deal)
}

func (h *DealHandler) GetDeals(c echo.Context) error {
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

	deals, err := h.Service.GetAll(page)
	if err != nil {
		log.Printf("Error retrieving deals: %v", err)
		return c.JSON(http.StatusInternalServerError, "Error retrieving deals")
	}

	return c.JSON(http.StatusOK, deals)
}

func (h *DealHandler) UpdateDeal(c echo.Context) error {
	var deal models.Deal
	if err := c.Bind(&deal); err != nil {
		log.Printf("Invalid input data: %v", err)
		return c.JSON(http.StatusBadRequest, "Invalid input data")
	}

	claims, ok := c.Get("userClaims").(*middlewares.Claims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, "You don't have rights")
	}

	updatedDeal, err := h.Service.Update(deal, claims)
	if err != nil {
		log.Printf("Error updating deal: %v", err)
		return c.JSON(http.StatusBadRequest, "Error updating deal")
	}

	return c.JSON(http.StatusOK, updatedDeal)
}

func (h *DealHandler) DeleteDeal(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid deal ID: %v", err)
		return c.JSON(http.StatusBadRequest, "Invalid deal ID")
	}

	claims, ok := c.Get("userClaims").(*middlewares.Claims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, "You don't have rights")
	}

	if err := h.Service.Delete(id, claims); err != nil {
		log.Printf("Error deleting deal: %v", err)
		return c.JSON(http.StatusNotFound, "Deal not found")
	}

	return c.NoContent(http.StatusOK)
}
