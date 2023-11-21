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

// PracticeDeleter отвечает за работу по удалению практических работ.
// Функции возвращают nil в случае успеха
type PracticeDeleter interface {
	// DeleteFile удаляет файл практической работы
	DeleteFile(id int) error
	// DeleteInfo удаляет информацию о практической работе
	DeleteInfo(id int) error
}
type Repository struct {
	PracticeSaver
	PracticeGetter
	PracticeDeleter
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		PracticeSaver:   NewPracticeRepository(db),
		PracticeGetter:  NewPracticeGetterRepository(db),
		PracticeDeleter: NewPracticeDeleterRepository(db),
	}
}
