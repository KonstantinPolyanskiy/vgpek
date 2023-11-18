package storage

import (
	"github.com/jmoiron/sqlx"
	"io"
	"log"
	"mime/multipart"
	"os"
	"practise-task-service/pkg/models"
)

const storagePath = "./pkg/storage/practiceStorageFolder/"

type PractiseRepository struct {
	db *sqlx.DB
}

func init() {
	err := os.Mkdir(storagePath, 0755)
	if err != nil {
		log.Printf("Ошибка в создании папки - %s\n", err)
	}
}

func (r *PractiseRepository) SaveMetadata(request models.UploadPracticeRequest, name string) (int, error) {
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

	err = tx.QueryRow(savePracticeInfoQuery, storagePath+name, "Холодов А.А.", request.Title, request.Theme, request.AcademicSubject).Scan(&practiceID)
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

func (r *PractiseRepository) RecordFile(practiceFile multipart.File, name string) error {

	dst, err := os.Create(storagePath + name)
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

func NewPracticeRepository(db *sqlx.DB) *PractiseRepository {
	return &PractiseRepository{
		db: db,
	}
}
