package usecase

import (
	"fmt"
	"github.com/essentialkaos/translit/v2"
	"math/rand"
	"mime/multipart"
	"path/filepath"
	"practise-task-service/internal/models"
	"practise-task-service/internal/storage"
)

type Saver interface {
	SaveMetadata(request models.UploadPracticeRequest, name string) (int, error)
	RecordFile(practiceFile multipart.File, name string) error
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

func NewUploadService(repo storage.PracticeSaver) *UploadService {
	return &UploadService{
		repo: repo,
	}
}

func (s *UploadService) Save(request models.UploadPracticeRequest, fh *multipart.FileHeader) (int, error) {
	name := translit.ICAO(fmt.Sprintf(
		"%s_%s_%s_%d_%s",
		request.AcademicSubject, request.Title, request.Theme, rand.Intn(1000), filepath.Ext(fh.Filename)))

	err := s.repo.RecordFile(request.File, name)
	if err != nil {
		return 0, err
	}

	id, err := s.repo.SaveMetadata(request, name)
	if err != nil {
		return 0, err
	}

	return id, nil
}
