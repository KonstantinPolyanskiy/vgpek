package usecase

import (
	"mod/internal/models"
	"mod/internal/storage"
)

type PracticeSaver interface {
	SaveMetadata(request models.UploadPracticeRequest) (int, error)
	RecordFile(request models.UploadPracticeRequest) error
}
type Service struct {
	PracticeSaver
}

func New(repo *storage.Repository) *Service {
	return &Service{
		PracticeSaver: NewUploadService(),
	}
}
