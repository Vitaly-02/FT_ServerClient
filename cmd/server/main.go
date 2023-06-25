package main

import (
	"FT_ServerClient/internal/server"
)

func main() {
	httpServer := server.New()
	httpServer.Start()
}
