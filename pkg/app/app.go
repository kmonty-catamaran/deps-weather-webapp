package app

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

type App struct {
	wg WeatherGetter
}

type WeatherGetter interface {
	GetWeather(ip string) (string, error)
}

func New(wg WeatherGetter) *App {
	return &App{wg: wg}
}

func (a *App) Handler() http.Handler {
	h := http.NewServeMux()
	h.HandleFunc("/", a.index)
	return h
}

func (a *App) index(w http.ResponseWriter, r *http.Request) {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		ip := net.ParseIP(r.RemoteAddr)
		if ip == nil {
			a.serverError(w, "invalid IP address %q", r.RemoteAddr)
			return
		}
		host = ip.String()
	}

	weather, err := a.wg.GetWeather(host)
	if err != nil {
		a.serverError(w, "ipw.GetWeather(): %v", err)
		return
	}
	log.Printf("Weather: %q\n", weather)

	fmt.Fprintf(w, "Weather for %q: %s", host, weather)
}

func (a *App) serverError(w http.ResponseWriter, f string, args ...any) {
	msg := fmt.Sprintf(f, args...)
	log.Println("Server error: " + msg)
	http.Error(w, msg, http.StatusInternalServerError)
}
