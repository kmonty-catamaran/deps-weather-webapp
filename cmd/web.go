package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/kmonty-catamaran/deps-weather-webapp/pkg/app"
)

const (
	defaultPort = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	a := app.New()

	if err := http.ListenAndServe(":"+port, a.Handler()); err != nil {
		log.Printf("Error: %s\n", err)
	}
}
