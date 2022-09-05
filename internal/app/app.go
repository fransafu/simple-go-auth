package app

import (
	"fmt"
	"net/http"

	"github.com/fransafu/simple-go-auth/internal/app/configuration"
	"github.com/fransafu/simple-go-auth/internal/app/database"
	"github.com/fransafu/simple-go-auth/internal/app/handlers"
	"github.com/fransafu/simple-go-auth/internal/app/router"
	"github.com/fransafu/simple-go-auth/internal/app/services"
	"github.com/go-chi/chi/v5"
)

func NewApp() {
	// Load configs
	configs, err := configuration.GetConfigs()
	if err != nil {
		panic(err)
	}

	// Load Web server
	r := chi.NewRouter()

	// Load Database
	database.Initialize()

	// Load Services
	services := services.NewServices()

	// Load Handlers
	handlers := handlers.NewHandlers(services)

	// Load Router
	router.LoadRoutes(r, handlers)

	err = http.ListenAndServe(fmt.Sprintf(":%s", configs.AppPort), r)
	if err != nil {
		panic(err)
	}
}
