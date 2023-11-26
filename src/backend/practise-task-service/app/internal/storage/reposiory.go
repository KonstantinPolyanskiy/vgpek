package storage

import (
	"github.com/jmoiron/sqlx"
	"log/slog"
	"mime/multipart"
	"practise-task-service/internal/models"
)

type PracticeSaver interface {
	SaveMetadata(request models.UploadPracticeRequest, name string) (int, error)
	RecordFile(practiceFile multipart.File, name string) error
}
type PracticeGetter interface {
	GetPracticeInfo(id int) (models.PracticeInfo, error)
	GetPracticeFile(id int) (models.PracticeFile, error)
	GetPracticeGroupInfo() (models.PracticesInfo, error)
	GetPracticeBySearch(title, subject string) (models.PracticesInfo, error)
	GetPracticePath(id int) (string, error)
}

// PracticeDeleter отвечает за работу по удалению практических работ.
// Функции возвращают nil в случае успеха
type PracticeDeleter interface {
	// DeleteFile удаляет файл практической работы
	DeleteFile(id int, deletedPath string) error
	// DeleteInfo удаляет информацию о практической работе
	DeleteInfo(id int) (string, error)
}
type Repository struct {
	PracticeSaver
	PracticeGetter
	PracticeDeleter
}

func New(db *sqlx.DB, savePath, deletePath string, Logger *slog.Logger) *Repository {
	Saver := NewPracticeRepository(db, savePath, Logger)
	Getter := NewPracticeGetterRepository(db, savePath, deletePath, Logger)
	Deleter := NewPracticeDeleterRepository(db, Getter, savePath, deletePath, Logger)

	return &Repository{
		PracticeSaver:   Saver,
		PracticeGetter:  Getter,
		PracticeDeleter: Deleter,
	}
}
