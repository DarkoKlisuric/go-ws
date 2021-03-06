package main

import (
	"net/http"

	"github.com/DarkoKlisuric/ws/internal/handlers"
	"github.com/bmizerany/pat"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	assets := http.FileServer(http.Dir("./assets/js/"))
	mux.Get("/assets/js/", http.StripPrefix("/assets/js", assets))

	return mux
}
