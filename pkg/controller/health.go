package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	Health struct {
		Message string `json:"message"`
	}
)

func GetHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, Health{"server OK"})
}
