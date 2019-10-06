package main

import "github.com/Pickausernaame/Ambitine_backend/server"

func main() {
	app := server.New()
	app.Run(":6666")
}
