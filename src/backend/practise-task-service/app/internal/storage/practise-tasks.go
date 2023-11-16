package storage

import (
	"github.com/jmoiron/sqlx"
	"mod/internal/models"
)

type PractiseRepository struct {
	db *sqlx.DB
}

func (r *PractiseRepository) SaveMetadata(request models.UploadPracticeRequest) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PractiseRepository) RecordFile(request models.UploadPracticeRequest) error {
	//TODO implement me
	panic("implement me")
}

func NewPracticeRepository(db *sqlx.DB) *PractiseRepository {
	return &PractiseRepository{
		db: db,
	}
}
