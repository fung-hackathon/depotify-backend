package server

import (
	"funhackathon2022-backend/pkg/controller"
	"funhackathon2022-backend/pkg/logger"
	"funhackathon2022-backend/pkg/models/firestore"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetRouter() *echo.Echo {
	e := echo.New()
	e.Use(logger.EchoLogger())

	if err := firestore.Initialize(); err != nil {
		logger.Log{
			Code:    http.StatusInternalServerError,
			Message: "failed to connect to firestore",
			Cause:   err,
		}.Err()
	}

	e.GET("/health", controller.GetHealth)
	e.POST("/user", controller.RegisterUser)
	e.GET("/score", controller.GetScore)
	e.POST("/score", controller.UpdateScore)
	e.GET("/emotion", controller.GetEmotion)
	e.POST("/emotion", controller.UpdateEmotion)
	return e
}

func Close() {
	firestore.Close()
}
