package storage

import (
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

type PracticeDeleterRepository struct {
	savePath, deletePath string
	db                   *sqlx.DB
}

func NewPracticeDeleterRepository(db *sqlx.DB, savePath, deletePath string) *PracticeDeleterRepository {
	return &PracticeDeleterRepository{
		db:         db,
		savePath:   savePath,
		deletePath: deletePath,
	}
}

func (r *PracticeDeleterRepository) DeleteFile(id int) error {
	path, err := getFilePath(id, r.db)
	if err != nil {
		return err
	}

	err = os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

func (r *PracticeDeleterRepository) DeleteInfo(id int) error {
	deletePracticeAccessQuery := `
	DELETE FROM practice_access
	WHERE practice_id=$1
`
	deletePracticeInfoQuery := `
	DELETE FROM practice_info
	WHERE id=$1
`

	_, err := r.db.Exec(deletePracticeAccessQuery, id)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(deletePracticeInfoQuery, id)
	if err != nil {
		return err
	}

	return nil
}

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
		log.Printf("Ошибка в получении пути практической работы - %s\v", err)
		return "", err
	}

	return path, nil
}
