package route

import "github.com/go-chi/chi"

func (s *Server) Routes() {
	s.router.Use(s.middleware.Recovery)

	s.router.Route("/user", func(router chi.Router) {
		router.With(s.middleware.Authorize).Get("/profile", s.handler.Profile)
	})

	s.router.Route("/storage", func(r chi.Router) {
		r.With(s.middleware.AuthorizeAdminOnly).Get("/list", s.handler.FetchStorage)
	})
}
