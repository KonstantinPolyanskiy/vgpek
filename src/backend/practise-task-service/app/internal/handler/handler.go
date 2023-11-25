package handler

import (
	"github.com/go-chi/chi/v5"
	"log/slog"
	"practise-task-service/internal/handler/mw"
	"practise-task-service/internal/usecase"
)

type Handler struct {
	service *usecase.Service
	logger  *slog.Logger
}

func New(service *usecase.Service, logger *slog.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) Init() *chi.Mux {
	r := chi.NewRouter()

	r.Use(mw.Logger(h.logger))

	r.Route("/api", func(r chi.Router) {
		r.Get("/practice-tasks", h.GetAllPracticeTask())
		r.Get("/practice-tasks/{id}", h.GetPractice())
		r.Get("/practice-tasks/search", h.SearchPractice())
		r.With(mw.ExtensionValidator(h.logger)).Post("/practice-tasks", h.UploadPractice())
		r.Delete("/practice-tasks/{id}", h.DeletePractice())
	})

	return r
}
