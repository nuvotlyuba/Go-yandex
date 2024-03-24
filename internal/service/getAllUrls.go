package service

import "github.com/nuvotlyuba/Go-yandex/internal/models"

func (s Service) GetAllUrls() []*models.URL, error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	storage := utils.SwitchStorage()
	logger.Info("get all URLs from", zap.String("storage", storage))
	switch storage {
	case "db":
		data, err := s.dbRepo.getAllURLs(ctx)
		if err != nil {
			return nil, err
		}
	case "file":
		data, err := s.fileRepo.ReadAllURLs()
		if err != nil {
			return nil, err
		}
	case "mem":
		data := s.memRepo.GetAllURLs()
	}

	return data, nil
}
