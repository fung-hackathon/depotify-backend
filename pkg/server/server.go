package server

import (
	"funhackathon2022-backend/pkg/controller"
	"funhackathon2022-backend/pkg/models/firestore"

	"github.com/labstack/echo/v4"
)

func GetRouter() *echo.Echo {
	e := echo.New()
	if err := firestore.Initialize(); err != nil {
		panic(err)
	}

	e.GET("/health", controller.GetHealth)
	e.POST("/user", controller.RegisterUser)
	e.GET("/score", controller.GetScore)
	e.POST("/score", controller.UpdateScore)
	return e
}

func Close() {
	firestore.Close()
}
