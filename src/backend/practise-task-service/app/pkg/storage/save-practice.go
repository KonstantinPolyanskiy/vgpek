package storage

import (
	"github.com/jmoiron/sqlx"
	"io"
	"log"
	"mime/multipart"
	"os"
	"practise-task-service/pkg/models"
)

const StoragePath = "./pkg/storage/practiceStorageFolder/"

type PractiseSaverRepository struct {
	db *sqlx.DB
}

func init() {
	err := os.Mkdir(StoragePath, 0755)
	if err != nil {
		log.Printf("Ошибка в создании папки - %s\n", err)
	}
}

func (r *PractiseSaverRepository) SaveMetadata(request models.UploadPracticeRequest, name string) (int, error) {
	var practiceID int

	savePracticeInfoQuery := `
	INSERT INTO practice_info 
	(relative_path, author, title, theme, academic_subject) 
	VALUES 
	($1, $2, $3, $4, $5)
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
		log.Printf("ошибка в запуске транзакции - %s\n", err)
		return 0, tx.Rollback()
	}

	err = tx.QueryRow(savePracticeInfoQuery, StoragePath+name, "Холодов А.А.", request.Title, request.Theme, request.AcademicSubject).Scan(&practiceID)
	if err != nil {
		log.Printf("ошибка в записи практической работы - %s\n", err)
		return 0, tx.Rollback()
	}

	for _, group := range request.AccessGroup {
		_, err := tx.Exec(saveAccessGroup, practiceID, group)
		if err != nil {
			log.Printf("ошибка в записи групп доступа - %s\n", err)
			return 0, tx.Rollback()
		}
	}

	return practiceID, tx.Commit()
}

func (r *PractiseSaverRepository) RecordFile(practiceFile multipart.File, name string) error {
	dst, err := os.Create(StoragePath + name)
	if err != nil {
		log.Println(err)
		return err
	}

	defer dst.Close()

	_, err = io.Copy(dst, practiceFile)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func NewPracticeRepository(db *sqlx.DB) *PractiseSaverRepository {
	return &PractiseSaverRepository{
		db: db,
	}
}
