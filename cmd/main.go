package main

import (
	"funhackathon2022-backend/pkg/server"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	e := server.GetRouter()
	e.Logger.Fatal(e.Start(":1323"))
}
