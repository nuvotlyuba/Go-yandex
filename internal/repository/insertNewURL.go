package repository

import (
	"github.com/nuvotlyuba/Go-yandex/configs"
	"github.com/nuvotlyuba/Go-yandex/internal/models"
)

func (r Repo) InsertNewURL(data *models.URL) error {

	if configs.FileStoragePath != "" {
		//сохраняем на диск в файле
		w, err := NewURLRecorder(configs.FileStoragePath)
		if err != nil {
			return err
		}
		err = w.WriteURL(data)
		defer w.Close()
		if err != nil {
			return err
		}

		return err
	}

	//берем значение из переменной
	DataURL = append(DataURL, data)
	return nil
}
