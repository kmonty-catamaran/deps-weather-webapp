package app

import (
	"net/http/httptest"
	"testing"
)

func TestAppIndex(t *testing.T) {
	a := New()
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	a.index(w, r)

	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}
	if w.Body.String() != "Hello, World!" {
		t.Errorf("Expected body 'Hello, World!', got '%s'", w.Body.String())
	}
}
