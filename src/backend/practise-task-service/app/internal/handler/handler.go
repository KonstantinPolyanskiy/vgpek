package handler

import "mod/internal/usecase"

type Handler struct {
	service usecase.Service
}

func New(service usecase.Service) *Handler {
	return &Handler{
		service: service,
	}
}
