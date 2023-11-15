package usecase

import (
	"mod/internal/models"
	"mod/internal/storage"
)

type PracticeSaver interface {
	Save(request models.UploadPracticeRequest) (int, error)
}
type Service struct {
	PracticeSaver
}

func New(repo *storage.Repository) *Service {
	return &Service{
		PracticeSaver: ,
	}
}
