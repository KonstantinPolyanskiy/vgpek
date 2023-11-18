package usecase

import (
	"mime/multipart"
	"practise-task-service/pkg/models"
	"practise-task-service/pkg/storage"
)

type PracticeSaver interface {
	Save(request models.UploadPracticeRequest, fh *multipart.FileHeader) (int, error)
}
type Service struct {
	PracticeSaver
}

func New(repo *storage.Repository) *Service {
	return &Service{
		PracticeSaver: NewUploadService(repo.PracticeSaver),
	}
}
