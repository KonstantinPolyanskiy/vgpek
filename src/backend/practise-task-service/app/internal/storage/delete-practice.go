package storage

import (
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"os"
	"path/filepath"
	log_err "practise-task-service/pkg/logger/error"
	"time"
)

type PracticeDeleterRepository struct {
	savePath, deletePath string
	db                   *sqlx.DB
	logger               *slog.Logger
	*PracticeGetterRepository
}

func NewPracticeDeleterRepository(db *sqlx.DB, getterRepository *PracticeGetterRepository, savePath, deletePath string, logger *slog.Logger) *PracticeDeleterRepository {
	return &PracticeDeleterRepository{
		db:                       db,
		savePath:                 savePath,
		deletePath:               deletePath,
		logger:                   logger,
		PracticeGetterRepository: getterRepository,
	}
}

func (r *PracticeDeleterRepository) DeleteFile(id int, deletedPath string) error {
	path, err := r.GetPracticePath(id)
	if err != nil {
		return err
	}
	name := filepath.Base(path)

	err = os.Rename(r.savePath+name, deletedPath)
	if err != nil {
		r.logger.Warn("ошибка удаления файла", log_err.Err(err))
		return err
	}

	return nil
}

func (r *PracticeDeleterRepository) DeleteInfo(id int) (string, error) {
	var deletedPath string

	path, err := r.GetPracticePath(id)
	if err != nil {
		return "", err
	}
	name := filepath.Base(path)

	deletePracticeInfoQuery := `
	UPDATE practice_info
	SET deleted_at=$1, relative_path=$2
	WHERE id=$3
	RETURNING relative_path
`

	row := r.db.QueryRow(deletePracticeInfoQuery, time.Now(), r.deletePath+name, id).Scan(&deletedPath)
	if errors.Is(row, pgx.ErrNoRows) {
		return "", nil
	}

	return deletedPath, nil
}

func (r *PracticeDeleterRepository) HardDeleteFile(name string) error {
	err := os.Remove(r.savePath + name)
	if err != nil {
		return err
	}

	return nil
}
