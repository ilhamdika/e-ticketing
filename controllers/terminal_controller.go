package controllers

import (
	"e-ticketing/config"
	"e-ticketing/models"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func CreateTerminal(c echo.Context) error {
	var terminal models.Terminal

	if err := c.Bind(&terminal); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request",
		})
	}

	if strings.TrimSpace(terminal.Name) == "" || strings.TrimSpace(terminal.Location) == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Field 'name' dan 'location' tidak boleh kosong.",
		})
	}

	userID := c.Get("user_id").(uint)
	terminal.CreatedBy = userID

	if err := config.DB.Create(&terminal).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Gagal menyimpan data terminal",
		})
	}

	return c.JSON(http.StatusOK, terminal)
}
