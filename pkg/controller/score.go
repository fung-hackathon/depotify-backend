package controller

import (
	"errors"
	"net/http"

	"funhackathon2022-backend/pkg/models"
	"funhackathon2022-backend/pkg/models/dto"
	"funhackathon2022-backend/pkg/models/firestore"

	"github.com/labstack/echo/v4"
)

var (
	ErrUnregisteredID = errors.New("unregistered user ID")
)

func QueryScore(userid dto.UserId) (int64, int, error) {

	obj, err := firestore.Get(userid)
	if err != nil || obj == nil {
		return 0, http.StatusBadRequest, ErrUnregisteredID
	}

	return obj["score"].(int64), http.StatusOK, nil
}

func GetScore(c echo.Context) error {
	var userid dto.UserId

	userid.UserId = c.QueryParam("userId")

	sc, status, err := QueryScore(userid)
	if err != nil {
		return c.JSON(status, dto.Error{status, err.Error()})
	}

	score := new(dto.Score)
	score.UserId = userid.UserId
	score.Score = sc

	return c.JSON(http.StatusOK, score)
}

func UpdateScore(c echo.Context) error {
	var coords dto.Coordinates

	err := c.Bind(&coords)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{http.StatusInternalServerError, err.Error()})
	}

	userid := dto.UserId{UserId: coords.UserId}

	currentScore, status, err := QueryScore(userid)
	if err != nil {
		return c.JSON(status, dto.Error{status, err.Error()})
	}

	incScore, err := models.CalculateScore(coords)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{http.StatusInternalServerError, err.Error()})
	}

	newScore := currentScore + incScore

	firestore.Set(userid, map[string]interface{}{
		"score": newScore,
	})

	score := new(dto.Score)
	score.UserId = coords.UserId

	score.Score = newScore

	return c.JSON(http.StatusOK, score)
}
