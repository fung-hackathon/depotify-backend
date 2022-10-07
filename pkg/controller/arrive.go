package controller

import (
	"errors"
	"funhackathon2022-backend/pkg/logger"
	"funhackathon2022-backend/pkg/models"
	"funhackathon2022-backend/pkg/models/dto"
	"funhackathon2022-backend/pkg/models/firestore"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrUnregisteredID = errors.New("unregistered user ID")
)

func queryScore(userid dto.UserId) (int64, int, error) {

	obj, err := firestore.Get(userid)
	if err != nil || obj == nil {
		return 0, http.StatusBadRequest, ErrUnregisteredID
	}

	return obj["score"].(int64), http.StatusOK, nil
}

func CheckArrival(c echo.Context) error {

	var userid dto.UserId

	userid.UserId = c.QueryParam("userId")

	currentScore, status, err := queryScore(userid)
	if err != nil {
		logger.Log{
			Code:    status,
			Message: err.Error(),
			Cause:   err,
		}.Warn()
		return c.JSON(status, dto.Error{status, err.Error()})
	}

	incScore, err := models.CalculateScore(dto.Coordinate{
		Longitude: c.QueryParam("olng"),
		Latitude:  c.QueryParam("olat")},
		dto.Coordinate{
			Longitude: c.QueryParam("dlng"),
			Latitude:  c.QueryParam("dlat"),
		})

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
	score.UserId = userid.UserId

	score.Score = newScore

	logger.Log{
		Code:    http.StatusOK,
		Message: "success to update score",
		Cause:   nil,
	}.Info()

	return c.JSON(http.StatusOK, score)
}
