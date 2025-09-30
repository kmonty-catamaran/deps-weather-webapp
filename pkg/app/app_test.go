package app

import (
	"fmt"
	"net"
	"net/http/httptest"
	"testing"
)

func TestAppIndex(t *testing.T) {
	weather := "stormy"
	a := New(&fakeWeatherGetter{weather: weather})
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
	want := fmt.Sprintf("Weather for %q: %s", host, weather)
	if got := w.Body.String(); got != want {
		t.Errorf("Incorrect body got %q, want %q", got, want)
	}
}

type fakeWeatherGetter struct {
	weather string
	err     error
}

func (f *fakeWeatherGetter) GetWeather(ip string) (string, error) {
	return f.weather, f.err
}
