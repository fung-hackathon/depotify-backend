package controller

import (
	"net/http"

	"funhackathon2022-backend/pkg/logger"
	"funhackathon2022-backend/pkg/models"
	"funhackathon2022-backend/pkg/models/dto"
	"funhackathon2022-backend/pkg/models/firestore"

	"github.com/labstack/echo/v4"
)

func RegisterUser(c echo.Context) error {
	userid := dto.UserId{UserId: models.GenerateUUID()}
	firestore.Set(userid, map[string]interface{}{
		"score": int64(0),
	})
	logger.Log{
		Code:    http.StatusOK,
		Message: "success to register user",
		Cause:   nil,
	}.Info()
	return c.JSON(http.StatusOK, userid)
}
