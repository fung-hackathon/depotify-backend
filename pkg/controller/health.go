package controller

import (
	"funhackathon2022-backend/pkg/models/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, dto.Health{"server OK"})
}
