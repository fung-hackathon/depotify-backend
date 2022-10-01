package controller

import (
	"net/http"

	"funhackathon2022-backend/pkg/models"

	"github.com/labstack/echo/v4"
)

type (
	UUID struct {
		UUID string `json:"uuid"`
	}
)

func GenerateUUID(c echo.Context) error {
	return c.JSON(http.StatusOK, UUID{UUID: models.GenerateUUID()})
}
