package service

import (
	"github.com/go-chi/chi/v5"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) RegisterHandlers(r chi.Router) {
	r.Post("/api/v1/calculate", s.CalculateHandler)
}
