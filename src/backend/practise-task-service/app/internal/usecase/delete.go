package usecase

import (
	"practise-task-service/internal/storage"
)

type DeleterService struct {
	repo storage.PracticeDeleter
}

func NewDeleterService(repo storage.PracticeDeleter) *DeleterService {
	return &DeleterService{
		repo: repo,
	}
}

func (s *DeleterService) Delete(id int) error {
	deletedPath, err := s.repo.DeleteInfo(id)
	if err != nil {
		return nil
	}

	err = s.repo.DeleteFile(id, deletedPath)
	if err != nil {
		return err
	}

	return nil
}
