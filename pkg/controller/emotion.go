package controller

import (
	"net/http"

	"funhackathon2022-backend/pkg/config"
	"funhackathon2022-backend/pkg/logger"
	"funhackathon2022-backend/pkg/models/dto"
	"funhackathon2022-backend/pkg/models/firestore"

	"github.com/labstack/echo/v4"
)

func QueryEmotion(userid dto.UserId) ([]string, int, error) {

	obj, err := firestore.Get(userid)
	if err != nil || obj == nil {
		return nil, http.StatusBadRequest, ErrUnregisteredID
	}

	eobj := obj["emotion"].([]interface{})

	ems := make([]string, len(eobj))
	for i := 0; i < len(eobj); i++ {
		ems[i] = eobj[i].(string)
	}

	return ems, http.StatusOK, nil
}

func GetEmotion(c echo.Context) error {
	var userid dto.UserId

	userid.UserId = c.QueryParam("userId")

	em, status, err := QueryEmotion(userid)
	if err != nil {
		logger.Log{
			Code:    status,
			Message: err.Error(),
			Cause:   err,
		}.Warn()
		return c.JSON(status, dto.Error{status, err.Error()})
	}

	emotion := new(dto.Emotion)
	emotion.UserId = userid.UserId
	emotion.Emotion = em

	return c.JSON(http.StatusOK, emotion)
}

func UpdateEmotion(c echo.Context) error {
	var userid dto.UserId

	userid.UserId = c.QueryParam("userId")

	currentEmotion, status, err := QueryEmotion(userid)
	if err != nil {
		logger.Log{
			Code:    status,
			Message: err.Error(),
			Cause:   err,
		}.Warn()
		return c.JSON(status, dto.Error{status, err.Error()})
	}

	newEmotion := append([]string{c.QueryParam("emotion")}, currentEmotion...)

	if len(newEmotion) > int(config.EMOTION_QUEUE_MAX_SIZE) {
		newEmotion = newEmotion[:config.EMOTION_QUEUE_MAX_SIZE]
	}

	emotion := new(dto.Emotion)
	emotion.UserId = userid.UserId

	emotion.Emotion = newEmotion

	firestore.Set(userid, map[string]interface{}{
		"emotion": newEmotion,
	})

	logger.Log{
		Code:    http.StatusOK,
		Message: "success to update score",
		Cause:   nil,
	}.Info()

	return c.JSON(http.StatusOK, emotion)
}
