package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Score struct {
		UserId string `json:"userid"`
		Score  int    `json:"score"`
	}

	Error struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	}

	Coordinate struct {
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lng"`
	}

	Coordinates struct {
		UserId      string       `json:"userid"`
		Coordinates []Coordinate `json:"coordinates"`
	}
)

func GetScore(c echo.Context) error {
	var userid UserId

	err := c.Bind(&userid)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Error{0, err.Error()})
	}

	// UUIDが登録済みであるかを判定 or 400

	score := new(Score)
	score.UserId = userid.UserId

	// スコアをscore.Scoreに入力

	return c.JSON(http.StatusOK, score)
}

func ReceiveScore(c echo.Context) error {
	var coords Coordinates

	err := c.Bind(&coords)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Error{0, err.Error()})
	}

	// UUIDが登録済みであるかを判定 or 400

	score := new(Score)
	score.UserId = coords.UserId

	// スコアを算出し、結果をscore.Scoreに入力

	return c.JSON(http.StatusOK, score)
}
