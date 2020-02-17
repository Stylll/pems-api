package main

import (
	"fmt"
	"os"

	Routes "github.com/Stylll/pems-api/src/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	startServer()
}

func startServer() {
	r := Routes.SetupRouter()
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	r.Run(port)
}
