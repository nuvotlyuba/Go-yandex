package repository

import (
	"fmt"

	"github.com/nuvotlyuba/Go-yandex/configs"
	"github.com/nuvotlyuba/Go-yandex/internal/models"
)

func (r Repo) InsertNewUrl(data *models.URL) error {

	if configs.FileStoragePath != "" {
		fmt.Println(configs.FileStoragePath, "path")
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
	DataUrl = append(DataUrl, data)
	fmt.Println(DataUrl, "DATA")
	return nil
}
