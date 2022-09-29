package server

import (
	"funhackathon2022-backend/pkg/controller"

	"github.com/labstack/echo"
)

func GetRouter() *echo.Echo {
	e := echo.New()
	e.GET("/health", controller.GetHealth)
	return e
}
