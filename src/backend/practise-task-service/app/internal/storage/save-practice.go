package storage

import (
	"github.com/jmoiron/sqlx"
	"io"
	"log/slog"
	"mime/multipart"
	"os"
	"practise-task-service/internal/models"
	log_err "practise-task-service/pkg/logger/error"
	"time"
)

type PractiseSaverRepository struct {
	savePath string
	db       *sqlx.DB
	logger   *slog.Logger
}

func (r *PractiseSaverRepository) SaveMetadata(request models.UploadPracticeRequest, name string) (int, error) {
	var practiceID int

	savePracticeInfoQuery := `
	INSERT INTO practice_info 
	(relative_path, author, title, theme, academic_subject, created_at) 
	VALUES 
	($1, $2, $3, $4, $5, $6)
	RETURNING id
`
	saveAccessGroup := `
	INSERT INTO practice_access 
	(practice_id, group_id) 
	VALUES 
	($1, (SELECT group_id FROM access_groups WHERE group_name=$2))
`
	tx, err := r.db.Begin()
	if err != nil {
		r.logger.Error("ошибка в запуске транзакции", log_err.Err(err))
		return 0, tx.Rollback()
	}

	err = tx.QueryRow(savePracticeInfoQuery,
		r.savePath+name, request.Author, request.Title, request.Theme, request.AcademicSubject, time.Now()).
		Scan(&practiceID)
	if err != nil {
		r.logger.Warn("ошибка в записи практической работы", log_err.Err(err))
		return 0, tx.Rollback()
	}

	for _, group := range request.AccessGroup {
		_, err := tx.Exec(saveAccessGroup, practiceID, group)
		if err != nil {
			r.logger.Warn("ошибка в записи групп доступа", log_err.Err(err))
			return 0, tx.Rollback()
		}
	}

	return practiceID, tx.Commit()
}

func (r *PractiseSaverRepository) RecordFile(practiceFile multipart.File, name string) error {
	dst, err := os.Create(r.savePath + name)
	if err != nil {
		r.logger.Warn("ошибка в создании дескриптора файла", log_err.Err(err))
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, practiceFile)
	if err != nil {
		r.logger.Warn("ошибка в записи файла практической работы", log_err.Err(err))
		return err
	}

	return nil
}

func NewPracticeRepository(db *sqlx.DB, savePath string, logger *slog.Logger) *PractiseSaverRepository {
	return &PractiseSaverRepository{
		db:       db,
		savePath: savePath,
		logger:   logger,
	}
}
