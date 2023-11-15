package usecase

import (
	"mod/internal/models"
	"os"
)

type UploadService struct {
	os.File
}

func NewUploadService() *UploadService {
	return &UploadService{}
}

func (s *Service) Save(request models.UploadPracticeRequest) (int, error) {
	return 1, nil
}
