package router

import (
	"github.com/NikhilSharma03/Yeager/internal/app"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Init(a *app.App) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	return r
}
