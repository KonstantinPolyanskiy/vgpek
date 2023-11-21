package usecase

import (
	"log"
	"practise-task-service/pkg/storage"
)

type Deleter interface {
	DeleteFile(id int) error
	DeleteInfo(id int) error
}

type DeleterService struct {
	repo storage.PracticeDeleter
}

func NewDeleterService(repo storage.PracticeDeleter) *DeleterService {
	return &DeleterService{
		repo: repo,
	}
}

func (s *DeleterService) Delete(id int) error {
	err := s.repo.DeleteFile(id)
	if err != nil {
		log.Printf("Ошибка в удалении файла - %s\n", err)
		return err
	}

	err = s.repo.DeleteInfo(id)
	if err != nil {
		log.Printf("Ошибка в удалении информации файла - %s\n", err)
		return err
	}

	return nil
}
