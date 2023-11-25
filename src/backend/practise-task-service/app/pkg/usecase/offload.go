package usecase

import (
	"practise-task-service/pkg/models"
	"practise-task-service/pkg/storage"
)

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
	info, err := s.repo.GetPracticeGroupInfo()
	if err != nil {
		return models.PracticesInfo{}, err
	}

	return info, nil
}

func (s *OffloadService) GetBySearch(title, subject string) (models.PracticesInfo, error) {
	practicesInfo, err := s.repo.GetPracticeBySearch(title, subject)
	if err != nil {
		return models.PracticesInfo{}, err
	}

	return practicesInfo, nil
}
