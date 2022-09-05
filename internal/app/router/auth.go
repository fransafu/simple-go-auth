package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func loadAuthRoutes(r chi.Router) {
	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {})
	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {})
	r.Post("/logout", func(w http.ResponseWriter, r *http.Request) {})
}
