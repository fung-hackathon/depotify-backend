package main

import (
	"funhackathon2022-backend/pkg/server"
	"os"
)

func main() {
	e := server.GetRouter()
	if os.Getenv("MODE") == "production" {
		e.HideBanner = true
		e.HidePort = true
	}

	e.Logger.Fatal(e.Start(":1323"))

	defer server.Close()
}
