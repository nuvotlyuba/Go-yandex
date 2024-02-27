package handlers

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/nuvotlyuba/Go-yandex/internal/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostUrlHandler(t *testing.T) {

	type want struct {
		contentType string
		statusCode  int
	}

	tests := []struct {
		name        string
		request     string
		url         string
		contentType string
		want        want
	}{
		{
			name:        "Success.Status code 201",
			request:     "/",
			contentType: "text/plain",
			url:         "https://yandex.ru",
			want: want{
				contentType: "text/plain",
				statusCode:  201,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, tt.request, strings.NewReader(tt.url))
			r.Header.Set("Content-Type", tt.contentType)
			s := new(Store)
			s.PostURLHandler(w, r)
			res := w.Result()

			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"), "Отличный от %s Content-Type", tt.want.contentType)
			assert.Equal(t, tt.want.statusCode, res.StatusCode, "Отличный от %d статус код", tt.want.statusCode)

			body, err := io.ReadAll(res.Body)
			require.NoError(t, err, "Ошибка чтения тела ответа")
			err = res.Body.Close()
			require.NoError(t, err)

			assert.NotEmpty(t, string(body), "Тело ответа пустое")
		})
	}
}

func testRequest(t *testing.T, ts *httptest.Server, method, path string) *http.Response {
	req, err := http.NewRequest(method, ts.URL+path, nil)
	require.NoError(t, err)
	resp, err := ts.Client().Do(req)
	require.NoError(t, err)
	resp.Body.Close()

	return resp
}

func TestGetUrlHandler(t *testing.T) {

	url := "https://yandex.ru"
	s:= new(services.Service)
	data, _ := s.CreateNewURL(url)
	fmt.Println(strings.Split(data.ShortURL, "/")[3], "ShortURL")

	type want struct {
		statusCode     int
		locationHeader string
	}

	tests := []struct {
		name        string
		request     string
		contentType string
		id          string
		want        want
	}{
		{
			name:    "Success.Status code 307",
			request:  "/" + strings.Split(data.ShortURL, "/")[3],
			want: want{
				statusCode:     200,
				locationHeader: url,
			},
		},
		{
			name:    "BadRequest.Status code 400",
			request: "/jhfybHYF",
			want: want{
				statusCode: 400,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := chi.NewRouter()
			ts := httptest.NewServer(BasicRouter(r))
			res := testRequest(t, ts, http.MethodGet, tt.request)
			res.Body.Close()

			assert.Equal(t, tt.want.statusCode, res.StatusCode, "Отличный от %d статус код", tt.want.statusCode)
			// assert.Equal(t, tt.want.locationHeader, res.Header.Get("Location"), "Отличный от %v заголовок Location", tt.want.locationHeader)

		})
	}
}

func TestPostURLJsonHandler(t *testing.T) {

	successBody := `{ "url": "https://yandex.ru" }`


	testCases := []struct {
		name         		string
		request      		string
		body         		string
		contentType  		string
		expectedCode 		int
		expectedContentType string
	}{
		{
			name:        		 "Success.Status code 201",
			request:     		 "/api/shorten",
			contentType: 		 "application/json",
			body:                successBody,
			expectedContentType: "application/json",
			expectedCode:        201,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, tt.request, strings.NewReader(tt.body))
			r.Header.Set("Content-Type", tt.contentType)
			s := new(Store)
			s.PostURLJsonHandler(w, r)
			res := w.Result()

			assert.Equal(t, tt.expectedContentType, res.Header.Get("Content-Type"), "Отличный от %s Content-Type", tt.expectedContentType)
			assert.Equal(t, tt.expectedCode, res.StatusCode, "Отличный от %d статус код", tt.expectedCode)

			_, err := io.ReadAll(res.Body)
			require.NoError(t, err, "Ошибка чтения тела ответа")
			err = res.Body.Close()
			require.NoError(t, err)
		})
	}
}

func TestGzipCompression( t *testing.T) {

	testCases := []struct {
		name         		string
		request      		string
		body         		string
		contentType  		string
		expectedCode 		int
		expectedContentType string
	}{
		{
			name:        		 "Send gzip.Success.Status code 201.URL:/",
			request:     		 "/",
			contentType: 		 "text/plain; charset=utf-8",
			body:                "https://yandex.ru",
			expectedContentType: "text/plain",
			expectedCode:        201,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T){
			buf := bytes.NewBuffer(nil)
			zb := gzip.NewWriter(buf)
			_, err := zb.Write([]byte(tt.body))
			require.NoError(t, err)
			err = zb.Close()
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, tt.request, buf)
			r.Header.Set("Content-Type", tt.contentType)
			r.Header.Set("Content-Encoding", "gzip, deflate, br")
			s := new(Store)
			var res *http.Response
			if tt.contentType == "text/plain; charset=utf-8" {
				s.PostURLHandler(w,r)
				res = w.Result()
			}
			assert.Equal(t, tt.expectedContentType, res.Header.Get("Content-Type"), "Отличный от %s Content-Type", tt.expectedContentType)
			assert.Equal(t, tt.expectedCode, res.StatusCode, "Отличный от %d статус код", tt.expectedCode)

			_, err = io.ReadAll(res.Body)
			require.NoError(t, err, "Ошибка чтения тела ответа")
			err = res.Body.Close()
			require.NoError(t, err)
		})
	}
}
