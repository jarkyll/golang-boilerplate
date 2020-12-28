package main

import (
	"awesomeProject/router"
)

func main() {
	server := router.NewRouter()

	server.Run("localhost:8080")
}
