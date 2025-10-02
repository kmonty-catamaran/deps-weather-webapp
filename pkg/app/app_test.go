package app

import (
	"fmt"
	"net"
	"net/http/httptest"
	"testing"

	ipweather "github.com/squee1945/deps-ip-weather"
)

func TestAppIndex(t *testing.T) {
	details := &ipweather.WeatherDetails{
		City:        "Mountain View",
		Region:      "CA",
		Country:     "US",
		Temperature: 20.0,
		Conditions:  "Partly Cloudy",
		Humidity:    60,
	}
	a := New(&fakeWeatherGetter{details: details})
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	a.index(w, r)

	if got, want := w.Code, 200; got != want {
		t.Errorf("Incorrect status code got %d, want %d", got, want)
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		t.Fatalf("net.SplitHostPort(): %v", err)
	}
	want := fmt.Sprintf("Weather for %q: %s", host, formatWeather(details))
	if got := w.Body.String(); got != want {
		t.Errorf("Incorrect body got %q, want %q", got, want)
	}
}

type fakeWeatherGetter struct {
	details *ipweather.WeatherDetails
	err     error
}

func (f *fakeWeatherGetter) GetWeather(ip string) (*ipweather.WeatherDetails, error) {
	return f.details, f.err
}
