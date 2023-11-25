package handler

import (
	"github.com/go-chi/chi/v5"
	"practise-task-service/pkg/handler/mw"
	"practise-task-service/pkg/usecase"
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
		r.With(mw.ExtensionValidator).Post("/practice-tasks", h.UploadPractice())
		r.Delete("/practice-tasks/{id}", h.DeletePractice())
	})

	return r
}
