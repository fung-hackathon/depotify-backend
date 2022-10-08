package config

import (
	"os"
)

var (
	YOLP_APPID                     string
	GOOGLE_APPLICATION_CREDENTIALS string
	ARRIVAL_REDIRECT_URL           string
	EMOTION_QUEUE_MAX_SIZE         uint
)

func init() {
	YOLP_APPID = os.Getenv("YOLP_APPID")
	GOOGLE_APPLICATION_CREDENTIALS = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	ARRIVAL_REDIRECT_URL = os.Getenv("ARRIVAL_REDIRECT_URL")
	EMOTION_QUEUE_MAX_SIZE = 30
}
