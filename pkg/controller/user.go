package controller

import (
	"net/http"

	"funhackathon2022-backend/pkg/models"

	"github.com/labstack/echo/v4"
)

type (
	UserId struct {
		UserId string `json:"userid"`
	}
)

func GenerateUUID(c echo.Context) error {
	return c.JSON(http.StatusOK, UserId{UserId: models.GenerateUUID()})
}
