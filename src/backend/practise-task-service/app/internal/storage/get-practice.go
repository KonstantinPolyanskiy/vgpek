package storage

import (
	"github.com/jmoiron/sqlx"
	"log"
	"log/slog"
	"os"
	"practise-task-service/internal/models"
	log_err "practise-task-service/pkg/logger/error"
)

type PracticeGetterRepository struct {
	savePath, deletePath string
	db                   *sqlx.DB
	logger               *slog.Logger
}

func NewPracticeGetterRepository(db *sqlx.DB, savePath, deletePath string, logger *slog.Logger) *PracticeGetterRepository {
	return &PracticeGetterRepository{
		db:         db,
		savePath:   savePath,
		deletePath: deletePath,
		logger:     logger,
	}
}

func (r *PracticeGetterRepository) GetPracticeInfo(id int) (models.PracticeInfo, error) {
	var info models.PracticeInfo

	getPracticeInfoQuery := `
	SELECT author, title, theme, academic_subject
	FROM practice_info
	WHERE id=$1 AND deleted_at IS NOT NULL
`
	err := r.db.Get(&info, getPracticeInfoQuery, id)
	if err != nil {
		log.Printf("Ошибка в получении информации о практической - %s\n", err)
		return models.PracticeInfo{}, err
	}

	return info, nil
}

func (r *PracticeGetterRepository) GetPracticeFile(id int) (models.PracticeFile, error) {
	var path string
	var practiceFile models.PracticeFile

	getPracticePathQuery := `
	SELECT relative_path 
	FROM practice_info 
	WHERE id=$1 AND deleted_at IS NOT NULL
`
	err := r.db.Get(&path, getPracticePathQuery, id)
	if err != nil {
		r.logger.Warn("ошибка в получении пути практической работы", log_err.Err(err))
		return models.PracticeFile{}, err
	}

	file, err := os.Open(path)
	if err != nil {
		r.logger.Warn("ошибка в открытии файла практической работы", log_err.Err(err))
		return models.PracticeFile{}, err
	}

	practiceFile.File = *file
	return practiceFile, nil
}

func (r *PracticeGetterRepository) GetPracticeGroupInfo() (models.PracticesInfo, error) {
	var practicesInfo models.PracticesInfo

	getPracticesInfoQuery := `
	SELECT author, title, theme, academic_subject
	FROM practice_info
	WHERE deleted_at IS NULL
	ORDER BY id
	LIMIT 50
`

	err := r.db.Select(&practicesInfo, getPracticesInfoQuery)
	if err != nil {
		r.logger.Warn("ошибка в получении информации о группе практических работ", log_err.Err(err))
		return models.PracticesInfo{}, err
	}

	return practicesInfo, nil
}

func (r *PracticeGetterRepository) GetPracticeBySearch(title, subject string) (models.PracticesInfo, error) {
	var practicesInfo models.PracticesInfo

	getPracticesQuery := `
	SELECT author, title, theme, academic_subject 
	FROM practice_info
	WHERE to_tsvector('russian', title) @@ to_tsquery('russian', $1)
	OR to_tsvector('russian', academic_subject) @@ to_tsquery('russian', $2)
`

	err := r.db.Select(&practicesInfo, getPracticesQuery, title, subject)
	if err != nil {
		r.logger.Warn("ошибка в получении информации по поиску о группе практических работ", log_err.Err(err))
		return models.PracticesInfo{}, err
	}

	return practicesInfo, nil
}
