package main

import "FT_ServerClient/internal/client"

func main() {
	httpClient := client.New()
	httpClient.Start()
}
