package server

import (
	"funhackathon2022-backend/pkg/controller"

	"github.com/labstack/echo/v4"
)

func GetRouter() *echo.Echo {
	e := echo.New()
	e.GET("/health", controller.GetHealth)
	e.POST("/user", controller.GenerateUUID)
	e.GET("/score", controller.GetScore)
	e.POST("/score", controller.ReceiveScore)
	return e
}
