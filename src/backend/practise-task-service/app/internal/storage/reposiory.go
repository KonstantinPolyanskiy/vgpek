package storage

import (
	"github.com/jmoiron/sqlx"
	"mod/internal/models"
)

type PracticeSaver interface {
	SaveMetadata(request models.UploadPracticeRequest) (int, error)
	RecordFile(request models.UploadPracticeRequest) error
}
type Repository struct {
	PracticeSaver
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		PracticeSaver: NewPracticeRepository(db),
	}
}
