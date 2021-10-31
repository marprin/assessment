package route

import "github.com/go-chi/chi"

func (s *Server) Routes() {
	s.router.Use(s.middleware.Recovery)

	s.router.Route("/user", func(router chi.Router) {
		router.With(s.middleware.Authorize).Get("/profile", s.handler.Profile)
	})
}
