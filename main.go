package main

import (
	"client/routes"
	"client/services"
	"fmt"
)

func main() {
	fmt.Println("Starting server at http://localhost:" + services.GetEnv("CLIENT_SERVER_PORT"))
	routes.HandleRequest()
}
