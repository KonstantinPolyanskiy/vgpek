package usecase

import (
	"mime/multipart"
	"practise-task-service/pkg/models"
	"practise-task-service/pkg/storage"
)

type PracticeSaver interface {
	Save(request models.UploadPracticeRequest, fh *multipart.FileHeader) (int, error)
}
type PracticeGetter interface {
	Get(id int) (models.PracticeResponse, error)
	GetGroup() (models.PracticesInfo, error)
}
type PracticeDeleter interface {
	Delete(id int) error
}

type Service struct {
	PracticeSaver
	PracticeGetter
	PracticeDeleter
}

func New(repo *storage.Repository) *Service {
	return &Service{
		PracticeSaver:   NewUploadService(repo.PracticeSaver),
		PracticeGetter:  NewOffloadService(repo.PracticeGetter),
		PracticeDeleter: NewDeleterService(repo.PracticeDeleter),
	}
}
