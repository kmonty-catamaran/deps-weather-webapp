package app

import "net/http"

type App struct {
}

func New() *App {
	return &App{}
}

func (a *App) Handler() http.Handler {
	h := http.NewServeMux()
	h.HandleFunc("/", a.index)
	return h
}

func (a *App) index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
