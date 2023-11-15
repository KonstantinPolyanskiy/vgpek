package handler

import (
	"github.com/go-chi/chi/v5"
	"mod/internal/usecase"
)

type Handler struct {
	service *usecase.Service
}

func New(service *usecase.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Init() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Get("/practice-tasks", h.GetAllPracticeTask())
		r.Get("/practice-tasks/{id}", h.GetPractice())
		r.Get("/practice-tasks/search?title={t}item={i}", h.SearchPractice())
		r.Post("practice-tasks", h.UploadPractice())
		r.Delete("practice-tasks/{id}", h.DeletePractice())
	})

	return r
}
