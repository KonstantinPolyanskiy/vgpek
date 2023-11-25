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
}

func NewPracticeDeleterRepository(db *sqlx.DB, savePath, deletePath string, logger *slog.Logger) *PracticeDeleterRepository {
	return &PracticeDeleterRepository{
		db:         db,
		savePath:   savePath,
		deletePath: deletePath,
		logger:     logger,
	}
}

func (r *PracticeDeleterRepository) DeleteFile(id int) error {
	path, err := getFilePath(id, r.db)
	if err != nil {
		r.logger.Warn("ошибка получения пути к файлу из базы данных", log_err.Err(err))
		return err
	}

	name := filepath.Base(path)

	err = os.Rename(path, r.deletePath+name)
	if err != nil {
		r.logger.Warn("ошибка удаления файла", log_err.Err(err))
		return err
	}

	return nil
}

func (r *PracticeDeleterRepository) DeleteInfo(id int) error {

	deletePracticeInfoQuery := `
	UPDATE practice_info
	SET deleted_at=$1
	WHERE id=$2
`

	_, err := r.db.Exec(deletePracticeInfoQuery, time.Now(), id)
	if err != nil {
		r.logger.Warn("ошибка в изменении статуса практической на <удаленный>", log_err.Err(err))
		return err
	}

	return nil
}

// Возвращает из базы данных путь к файлу (даже если он помечен как удаленный)
func getFilePath(id int, db *sqlx.DB) (string, error) {
	var path string

	getPathQuery := `
	SELECT relative_path 
	FROM practice_info
	WHERE id=$1
`

	err := db.Get(&path, getPathQuery, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", nil
	}
	if err != nil {
		return "", err
	}

	return path, nil
}
