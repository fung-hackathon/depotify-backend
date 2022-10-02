package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

var (
	ErrGodotenv = errors.New("faild to load .env file")
)

var (
	YOLP_APPID                     string
	GOOGLE_APPLICATION_CREDENTIALS string
)

func init() {
	err := godotenv.Load("./../../.env")
	if err != nil {
		panic(err)
	}
	YOLP_APPID = os.Getenv("YOLP_APPID")
	GOOGLE_APPLICATION_CREDENTIALS = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
}
