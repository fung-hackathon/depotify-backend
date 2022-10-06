package config

import (
	"os"
)

var (
	YOLP_APPID                     string
	GOOGLE_APPLICATION_CREDENTIALS string
)

func init() {
	YOLP_APPID = os.Getenv("YOLP_APPID")
	GOOGLE_APPLICATION_CREDENTIALS = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
}
