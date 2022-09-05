package router

import (
	"github.com/fransafu/simple-go-auth/internal/app/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func LoadRoutes(router *chi.Mux, handlers *handlers.Handlers) {
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Route("/api", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", handlers.AuthHandler.Login)
			r.Post("/register", handlers.AuthHandler.Register)
			r.Post("/logout", handlers.AuthHandler.Logout)
		})

		r.With(authMiddleware).Route("/users", func(r chi.Router) {
			r.Get("/", handlers.UserHandler.GetAll)
			r.Get("/{id}", handlers.UserHandler.GetByID)
		})
	})
}
