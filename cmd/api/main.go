package main

import "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/cmd/api/server"

func main() {
	_ = server.New().Run(":8080")
}
