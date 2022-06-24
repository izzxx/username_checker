package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/izzxx/username_checker/components/controller"
)

type Dependency struct {
	App *chi.Mux
}

func (deps *Dependency) Check() {
	deps.App.Route("/api/check", func(r chi.Router) {
		r.Get("/facebook/{username}", controller.UsernameChecker)
		r.Get("/instagram/{username}", controller.UsernameChecker)
		r.Get("/github/{username}", controller.UsernameChecker)
	})
}
