package controller

import (
	"net/http"

	"funhackathon2022-backend/pkg/logger"
	"funhackathon2022-backend/pkg/models/dto"

	"github.com/labstack/echo/v4"
)

func GetScore(c echo.Context) error {
	var userid dto.UserId

	userid.UserId = c.QueryParam("userId")

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

/*
func UpdateScore(c echo.Context) error {
	var coords dto.Coordinates

	err := c.Bind(&coords)

	if err != nil {
		logger.Log{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Cause:   err,
		}.Err()
		return c.JSON(http.StatusInternalServerError, dto.Error{http.StatusInternalServerError, err.Error()})
	}

	userid := dto.UserId{UserId: coords.UserId}

	currentScore, status, err := QueryScore(userid)
	if err != nil {
		logger.Log{
			Code:    status,
			Message: err.Error(),
			Cause:   err,
		}.Warn()
		return c.JSON(status, dto.Error{status, err.Error()})
	}

	incScore, err := models.CalculateScore(coords)
	if err != nil {
		logger.Log{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Cause:   err,
		}.Err()
		return c.JSON(http.StatusInternalServerError, dto.Error{http.StatusInternalServerError, err.Error()})
	}

	newScore := currentScore + incScore

	firestore.Update(userid, "score", newScore)

	score := new(dto.Score)
	score.UserId = coords.UserId

	score.Score = newScore

	logger.Log{
		Code:    http.StatusOK,
		Message: "success to update score",
		Cause:   nil,
	}.Info()

	return c.JSON(http.StatusOK, score)
}

*/
