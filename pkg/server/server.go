package server

import (
	"funhackathon2022-backend/pkg/controller"
	"funhackathon2022-backend/pkg/logger"
	"funhackathon2022-backend/pkg/models/firestore"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetRouter() *echo.Echo {
	e := echo.New()
	e.Use(logger.EchoLogger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("FRONT_ORIGIN")},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

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
	e.GET("/emotion", controller.GetEmotion)
	e.POST("/emotion", controller.UpdateEmotion)
	e.GET("/arrive", controller.CheckArrival)
	return e
}

func Close() {
	firestore.Close()
}
