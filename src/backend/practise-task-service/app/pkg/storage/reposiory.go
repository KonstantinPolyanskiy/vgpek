package storage

import (
	"github.com/jmoiron/sqlx"
	"mime/multipart"
	"practise-task-service/pkg/models"
)

type PracticeSaver interface {
	SaveMetadata(request models.UploadPracticeRequest, name string) (int, error)
	RecordFile(practiceFile multipart.File, name string) error
}
type PracticeGetter interface {
	GetPracticeInfo(id int) (models.PracticeInfo, error)
	GetPracticeFile(id int) (models.PracticeFile, error)
}
type Repository struct {
	PracticeSaver
	PracticeGetter
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		PracticeSaver:  NewPracticeRepository(db),
		PracticeGetter: NewPracticeGetterRepository(db),
	}
}
