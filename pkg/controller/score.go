package controller

import (
	"errors"
	"net/http"

	"funhackathon2022-backend/pkg/models"
	"funhackathon2022-backend/pkg/models/firestore"

	"github.com/labstack/echo/v4"
)

var (
	ErrUnregisteredID = errors.New("unregistered user ID")
)

func QueryScore(userid models.UserId) (int64, int, error) {

	obj, err := firestore.Get(userid)
	if err != nil || obj == nil {
		return 0, http.StatusBadRequest, ErrUnregisteredID
	}

	return obj["score"].(int64), http.StatusOK, nil
}

func GetScore(c echo.Context) error {
	var userid models.UserId

	userid.UserId = c.QueryParam("userId")

	sc, status, err := QueryScore(userid)
	if err != nil {
		return c.JSON(status, models.Error{status, err.Error()})
	}

	score := new(models.Score)
	score.UserId = userid.UserId
	score.Score = sc

	return c.JSON(http.StatusOK, score)
}

func UpdateScore(c echo.Context) error {
	var coords models.Coordinates

	err := c.Bind(&coords)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Error{http.StatusInternalServerError, err.Error()})
	}

	// UUIDが登録済みであるかを判定 or 400
	sc, status, err := QueryScore(models.UserId{UserId: coords.UserId})
	if err != nil {
		return c.JSON(status, models.Error{status, err.Error()})
	}

	score := new(models.Score)
	score.UserId = coords.UserId
	score.Score = sc + models.CalculateScore(coords)

	return c.JSON(http.StatusOK, score)
}
