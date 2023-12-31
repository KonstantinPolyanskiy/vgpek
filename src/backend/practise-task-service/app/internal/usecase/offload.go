package usecase

import (
	"errors"
	"practise-task-service/internal/models"
	"practise-task-service/internal/storage"
)

var ErrNoResult = errors.New("нет результата")

type OffloadService struct {
	repo storage.PracticeGetter
}

func NewOffloadService(repo storage.PracticeGetter) *OffloadService {
	return &OffloadService{
		repo: repo,
	}
}

func (s *OffloadService) Get(id int) (models.PracticeResponse, error) {
	info, err := s.repo.GetPracticeInfo(id)
	if err != nil {
		return models.PracticeResponse{}, err
	}

	if len(info.Author) == 0 {
		return models.PracticeResponse{}, ErrNoResult
	}

	file, err := s.repo.GetPracticeFile(id)
	if err != nil {
		return models.PracticeResponse{}, err
	}

	return models.PracticeResponse{
		PracticeInfo: info,
		PracticeFile: file,
	}, nil
}

func (s *OffloadService) GetGroup() (models.PracticesInfo, error) {
	practicesInfo, err := s.repo.GetPracticeGroupInfo()
	if err != nil {
		return models.PracticesInfo{}, err
	}

	if len(practicesInfo) == 0 {
		return models.PracticesInfo{}, ErrNoResult
	}

	return practicesInfo, nil
}

func (s *OffloadService) GetBySearch(title, subject string) (models.PracticesInfo, error) {
	practicesInfo, err := s.repo.GetPracticeBySearch(title, subject)
	if err != nil {
		return models.PracticesInfo{}, err
	}

	if len(practicesInfo) == 0 {
		return models.PracticesInfo{}, ErrNoResult
	}

	return practicesInfo, nil
}
