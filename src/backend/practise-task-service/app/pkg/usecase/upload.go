package usecase

import (
	"fmt"
	"github.com/essentialkaos/translit/v2"
	"log"
	"mime/multipart"
	"path/filepath"
	"practise-task-service/pkg/models"
	"practise-task-service/pkg/storage"
	"strings"
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
	//TODO: сохранить файл в директорию
	//TODO: сохранить данные о файле в репозитории

	name := translit.Scientific(fmt.Sprintf("%s %s %s", request.AcademicSubject, request.Title, request.Theme))
	name = strings.Replace(name, " ", "_", -1)
	name = fmt.Sprintf("%s%s", name, filepath.Ext(fh.Filename))

	err := s.repo.RecordFile(request.File, name)
	if err != nil {
		log.Printf("Ошибка в записи файла - %s\n", err)
		return 0, err
	}

	id, err := s.repo.SaveMetadata(request, name)

	return id, nil
}
