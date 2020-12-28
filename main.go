package main

import (
	"golang-boilerplate/router"
)

func main() {
	server := router.NewRouter()

	server.Run("localhost:8080")
}
