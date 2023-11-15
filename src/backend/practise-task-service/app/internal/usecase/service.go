package usecase

import "mod/internal/storage"

type Service struct {
}

func New(repo storage.Repository) *Service {
	return &Service{}
}
