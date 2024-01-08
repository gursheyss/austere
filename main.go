package main

import (
	"austere/internal/server"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, loading environment variables from OS")
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			os.Setenv(pair[0], pair[1])
		}
	}
	server.StartServer()
}
