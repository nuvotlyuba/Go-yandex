package store

//куда лучше поместить данный файл???
import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
	"github.com/nuvotlyuba/Go-yandex/internal/models"
)

type URLRecorder struct {
	file   *os.File
	writer *bufio.Writer
}

func NewURLRecorder(filename string) (*URLRecorder, error) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return &URLRecorder{
		file:   file,
		writer: bufio.NewWriter(file),
	}, nil
}

func (w *URLRecorder) Close() error {
	return w.file.Close()
}

func (w *URLRecorder) WriteURL(url *models.URL) error {
	data, err := json.Marshal(&url)
	if err != nil {
		return err
	}

	if _, err = w.writer.Write(data); err != nil {
		return err
	}

	if err := w.writer.WriteByte('\n'); err != nil {
		return err
	}

	return w.writer.Flush()
}

type URLScanner struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewURLScanner(filename string) (*URLScanner, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_APPEND|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return &URLScanner{
		file:    file,
		scanner: bufio.NewScanner(file),
	}, nil
}

func (s *URLScanner) Split() {
	s.scanner.Split(bufio.ScanLines)
}

func (s *URLScanner) ScanURL(shortenURL string) (*models.URL, error) {
	url := models.URL{}
	var d models.URL

	for s.scanner.Scan() {
		data := s.scanner.Bytes()
		err := json.Unmarshal(data, &url)
		if err != nil {
			return nil, err
		}
		if url.ShortURL == shortenURL {
			d = url
			break
		}
	}

	if err := s.scanner.Err(); err != nil {
		return nil, err
	}
	return &d, nil

}

func (s *URLScanner) ScanAllURLs() ([]models.URL, error) {
	urls := make([]models.URL, 0)
	var url models.URL

	for s.scanner.Scan() {
		data := s.scanner.Bytes()
		err := json.Unmarshal(data, &url)
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	if err := s.scanner.Err(); err != nil {
		return nil, err
	}
	return urls, nil

}

type FileRepo interface {
	WriteNewUrl(data *models.URL) error
	ReadURL(shortURL string) (*models.URL, error)
	WriteBatchURL(data []*models.URL) error
}

type FileRepository struct {
	FileStoragePath string
}

func (r *FileRepository) WriteNewURL(data *models.URL) error {
	w, err := NewURLRecorder(r.FileStoragePath)
	if err != nil {
		return fmt.Errorf("error in FileRepository: WriteNewURL -> %v", err)
	}
	err = w.WriteURL(data)
	if err != nil {
		return fmt.Errorf("error in FileRepository: WriteNewURL -> %v", err)
	}

	return nil

}

func (r *FileRepository) ReadURL(shortenURL string) (string, error) {
	rr, err := NewURLScanner(r.FileStoragePath)
	if err != nil {
		return "", fmt.Errorf("error in FileRepository: ReadURL -> %v", err)
	}
	rr.Split()
	data, err := rr.ScanURL(shortenURL)
	if err != nil {
		return "", fmt.Errorf("error in FileRepository: ReadURL -> %v", err)
	}

	return data.OriginalURL, nil
}

func (r *FileRepository) WriteBatchURL(data []*models.URL) error {
	w, err := NewURLRecorder(r.FileStoragePath)
	if err != nil {
		return err
	}
	for _, v := range data {
		err = w.WriteURL(v)
		if err != nil {
			return fmt.Errorf("error in FileRepository: WriteBatchURL -> %v", err)
		}
	}

	return nil
}

func (r *FileRepository) ReadAllURLs() (*[]models.URL, error) {
	rr, err := NewURLScanner(r.FileStoragePath)
	if err != nil {
		logger.Debug("error in FileRepository: ReadAllURLs.NewURLScanner ->", err)
		return nil, ErrNoContent
	}
	rr.Split()
	data, err := rr.ScanAllURLs()
	if err != nil {
		return nil, fmt.Errorf("error in FileRepository: ReadAllURLs.ScanAllURLs -> %v", err)
	}

	return &data, nil
}
