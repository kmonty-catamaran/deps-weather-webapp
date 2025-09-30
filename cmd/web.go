package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/kmonty-catamaran/deps-weather-webapp/pkg/app"
	ipweather "github.com/squee1945/deps-ip-weather"
)

const (
	defaultPort = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	ipw, err := ipweather.New()
	if err != nil {
		exit("ipweather.New(): %v", err)
	}

	a := app.New(ipw)

	if err := http.ListenAndServe(":"+port, a.Handler()); err != nil {
		exit("http.ListenAndServe(): %v", err)
	}
}

func exit(f string, args ...any) {
	log.Printf("Error: "+f+"\n", args...)
	os.Exit(1)
}
