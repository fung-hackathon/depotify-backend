package controller

import (
	"net/http"

	"funhackathon2022-backend/pkg/config"
	"funhackathon2022-backend/pkg/logger"
	"funhackathon2022-backend/pkg/models/dto"
	"funhackathon2022-backend/pkg/models/firestore"

	"github.com/labstack/echo/v4"
)

func queryEmotion(userid dto.UserId) ([]string, int, error) {

	obj, err := firestore.Get(userid)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if obj == nil {
		return nil, http.StatusBadRequest, ErrUnregisteredID
	}

	if obj["emotion"] != nil {
		return nil, http.StatusInternalServerError, err
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

	em, status, err := queryEmotion(userid)
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

	var ems dto.Emotion

	err := c.Bind(&ems)
	if err != nil {
		logger.Log{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Cause:   err,
		}.Err()
		return c.JSON(http.StatusInternalServerError, dto.Error{http.StatusInternalServerError, err.Error()})
	}

	userid := dto.UserId{UserId: ems.UserId}

	currentEmotion, status, err := queryEmotion(userid)
	if err != nil {
		logger.Log{
			Code:    status,
			Message: err.Error(),
			Cause:   err,
		}.Warn()
		return c.JSON(status, dto.Error{status, err.Error()})
	}

	newEmotion := append(ems.Emotion, currentEmotion...)

	if len(newEmotion) > int(config.EMOTION_QUEUE_MAX_SIZE) {
		newEmotion = newEmotion[:config.EMOTION_QUEUE_MAX_SIZE]
	}

	emotion := new(dto.Emotion)
	emotion.UserId = ems.UserId

	emotion.Emotion = newEmotion

	err = firestore.Update(userid, "emotion", newEmotion)
	if err != nil {
		logger.Log{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Cause:   err,
		}.Err()
		return c.JSON(http.StatusInternalServerError, dto.Error{http.StatusInternalServerError, err.Error()})
	}

	logger.Log{
		Code:    http.StatusOK,
		Message: "success to update score",
		Cause:   nil,
	}.Info()

	return c.JSON(http.StatusOK, emotion)
}
