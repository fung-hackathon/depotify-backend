package main

import (
	"funhackathon2022-backend/pkg/server"
)

func main() {
	e := server.GetRouter()
	e.Logger.Fatal(e.Start(":1323"))
	defer server.Close()
}
