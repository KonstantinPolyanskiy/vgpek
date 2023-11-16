package usecase

import (
	"mod/internal/models"
	"mod/internal/storage"
)

type Saver interface {
	SaveMetadata(request models.UploadPracticeRequest) (int, error)
	RecordFile(request models.UploadPracticeRequest) error
}

type PracticeMetadata struct {
	Title           string   `json:"title"`
	Theme           string   `json:"theme"`
	AcademicSubject string   `json:"academicSubject"`
	AccessGroup     []string `json:"accessGroup"`
}

type UploadService struct {
	repo storage.PracticeSaver
}

func NewUploadService() *UploadService {
	return &UploadService{}
}

func (s *UploadService) Save(request models.UploadPracticeRequest) (int, error) {
	//TODO: сохранить файл в директорию
	//TODO: сохранить данные о файле в репозитории
	err := s.repo.RecordFile(request)
	if err != nil {
		return 0, err
	}
	id, err := s.repo.SaveMetadata(request)
	if err != nil {
		return 0, err
	}

	return id, nil
}
