package controller

import (
	"net/http"

	"funhackathon2022-backend/pkg/logger"
	"funhackathon2022-backend/pkg/models/dto"

	"github.com/labstack/echo/v4"
)

func GetScore(c echo.Context) error {
	var userid dto.UserId

	userid.UserId = c.QueryParam("userid")

	sc, status, err := queryScore(userid)
	if err != nil {
		logger.Log{
			Code:    status,
			Message: err.Error(),
			Cause:   err,
		}.Warn()
		return c.JSON(status, dto.Error{status, err.Error()})
	}

	score := new(dto.Score)
	score.UserId = userid.UserId
	score.Score = sc

	logger.Log{
		Code:    status,
		Message: "success to get score",
		Cause:   nil,
	}.Info()
	return c.JSON(http.StatusOK, score)
}
