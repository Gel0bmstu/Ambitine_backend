package main

import "./server"

func main() {
	app := server.New()
	app.Run(":6666")
}
