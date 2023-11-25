package usecase

import (
	"practise-task-service/internal/storage"
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
		return err
	}

	err = s.repo.DeleteInfo(id)
	if err != nil {
		return err
	}

	return nil
}
