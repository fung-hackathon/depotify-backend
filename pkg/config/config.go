package config

import (
	"os"
)

var (
	YOLP_APPID                     string
	GOOGLE_APPLICATION_CREDENTIALS string
	EMOTION_QUEUE_MAX_SIZE         uint
)

func init() {
	YOLP_APPID = os.Getenv("YOLP_APPID")
	GOOGLE_APPLICATION_CREDENTIALS = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	EMOTION_QUEUE_MAX_SIZE = 30
}
