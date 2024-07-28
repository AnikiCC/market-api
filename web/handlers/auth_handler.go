package handlers

import (
	"net/http"

	"market/internal/database/models"
	"market/internal/services"
	"market/web/handlers/middlewares"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	Service services.UserService
}

func (h *AuthHandler) Login(c echo.Context) error {
	var loginUser models.LoginUser

	if err := c.Bind(&loginUser); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := h.Service.Authenticate(loginUser.Username, loginUser.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid username or password")
	}

	tokenString, err := middlewares.GenerateJWT(user.Id, user.Username, false)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	refreshToken, err := middlewares.GenerateJWT(user.Id, user.Username, true)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token":         tokenString,
		"refresh_token": refreshToken,
	})
}

func (h *AuthHandler) RefreshToken(c echo.Context) error {
	var tokenReq struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.Bind(&tokenReq); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	claims, err := middlewares.GetValidatedClaims(tokenReq.RefreshToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid refresh token")
	}

	if !claims.Refresh {
		return c.JSON(http.StatusUnauthorized, "Invalid refresh token")
	}

	newToken, err := middlewares.GenerateJWT(claims.UserId, claims.Username, false)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not generate new token")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": newToken,
	})
}
