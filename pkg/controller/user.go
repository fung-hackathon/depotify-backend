package controller

import (
	"net/http"

	"funhackathon2022-backend/pkg/models"

	"github.com/labstack/echo/v4"
)

func GenerateUUID(c echo.Context) error {
	return c.JSON(http.StatusOK, models.UserId{UserId: models.GenerateUUID()})
}
