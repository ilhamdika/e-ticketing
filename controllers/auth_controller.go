package controllers

import (
	"e-ticketing/config"
	"e-ticketing/middleware"
	"e-ticketing/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request"})
	}

	var user models.User
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid credentials"})
	}

	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}
