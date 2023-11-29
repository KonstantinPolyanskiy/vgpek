package usecase

import (
	"fmt"
	"github.com/essentialkaos/translit/v2"
	"math/rand"
	"mime/multipart"
	"path/filepath"
	"practise-task-service/internal/models"
	"practise-task-service/internal/storage"
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
	saver   storage.PracticeSaver
	deleter storage.PracticeDeleter
}

func NewUploadService(saver storage.PracticeSaver, deleter storage.PracticeDeleter) *UploadService {
	return &UploadService{
		saver:   saver,
		deleter: deleter,
	}
}

func (s *UploadService) Save(request models.UploadPracticeRequest, fh *multipart.FileHeader) (int, error) {
	name := translit.ICAO(fmt.Sprintf("%s %s %s %d",
		request.AcademicSubject, request.Title, request.Theme, rand.Intn(10000)))
	name = strings.Replace(name, " ", "_", -1)
	name = fmt.Sprintf("%s%s", name, filepath.Ext(fh.Filename))

	err := s.saver.RecordFile(request.File, name)
	if err != nil {
		return 0, err
	}

	id, err := s.saver.SaveMetadata(request, name)
	if err != nil {
		err = s.deleter.HardDeleteFile(name)
		return 0, err
	}

	return id, nil
}
