package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/nuvotlyuba/Go-yandex/internal/repository"
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
			contentType: "text/plain; charset=utf-8",
			url:         "https://yandex.ru",
			want: want{
				contentType: "text/plain",
				statusCode:  201,
			},
		},
		{
			name:        "Unsupported Media Type.Status code 415",
			request:     "/",
			contentType: "application/json",
			url:         "https://yandex.ru",
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  415,
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
	r := new(repository.Repo)
	id := r.CreateNewID(url)

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
			request: "/" + id,
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
	errorBody := "https://yandex.ru"

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
		{
			name:                "Unsupported Media Type.Status code 415",
			request:             "/api/shorten",
			contentType:         "text/plain",
			body:                errorBody,
			expectedContentType: "",
			expectedCode:        415,
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
