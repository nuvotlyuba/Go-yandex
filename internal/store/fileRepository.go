package store

//куда лучше поместить данный файл???
import (
	"bufio"
	"encoding/json"
	"os"

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

type FileRepo interface {
	WriteNewUrl(data *models.URL) error
	ReadURL(shortURL string) (*models.URL, error)
}

type FileRepository struct {
	FileStoragePath string
}

func (r *FileRepository) WriteNewURL(data *models.URL) error {

	w, err := NewURLRecorder(r.FileStoragePath)
	if err != nil {
		return err
	}
	err = w.WriteURL(data)
	if err != nil {
		return err
	}

	return nil

}

func (r *FileRepository) ReadURL(shortenURL string) (*models.URL, error) {
	rr, err := NewURLScanner(r.FileStoragePath)
	if err != nil {
		return nil, err
	}
	rr.Split()
	data, err := rr.ScanURL(shortenURL)
	if err != nil {
		return nil, err
	}

	return data, nil
}
