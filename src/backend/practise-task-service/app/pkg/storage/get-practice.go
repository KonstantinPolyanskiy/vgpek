package storage

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"practise-task-service/pkg/models"
)

type PracticeGetterRepository struct {
	savePath, deletePath string
	db                   *sqlx.DB
}

func NewPracticeGetterRepository(db *sqlx.DB, savePath, deletePath string) *PracticeGetterRepository {
	return &PracticeGetterRepository{
		db:         db,
		savePath:   savePath,
		deletePath: deletePath,
	}
}

func (r *PracticeGetterRepository) GetPracticeInfo(id int) (models.PracticeInfo, error) {
	var info models.PracticeInfo

	getPracticeInfoQuery := `
	SELECT author, title, theme, academic_subject
	FROM practice_info
	WHERE id=$1
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
	SELECT relative_path FROM practice_info WHERE id=$1
`
	err := r.db.Get(&path, getPracticePathQuery, id)
	if err != nil {
		log.Printf("Ошибка в получении пути к практической работе - %s\n", err)
		return models.PracticeFile{}, err
	}

	file, err := os.Open(path)
	if err != nil {
		log.Printf("Ошибка в открытии файла - %s\n", err)
		return models.PracticeFile{}, err
	}

	practiceFile.File = *file
	return practiceFile, nil
}
