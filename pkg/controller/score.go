package controller

import (
	"net/http"

	"funhackathon2022-backend/pkg/models"

	"github.com/labstack/echo/v4"
)

func GetScore(c echo.Context) error {
	var userid models.UserId

	err := c.Bind(&userid)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Error{0, err.Error()})
	}

	// UUIDが登録済みであるかを判定 or 400

	score := new(models.Score)
	score.UserId = userid.UserId

	// スコアをscore.Scoreに入力

	return c.JSON(http.StatusOK, score)
}

func ReceiveScore(c echo.Context) error {
	var coords models.Coordinates

	err := c.Bind(&coords)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Error{0, err.Error()})
	}

	// UUIDが登録済みであるかを判定 or 400

	score := new(models.Score)
	score.UserId = coords.UserId

	// スコアを算出し、結果をscore.Scoreに入力

	return c.JSON(http.StatusOK, score)
}
