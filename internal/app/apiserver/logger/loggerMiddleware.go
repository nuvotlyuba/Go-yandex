package logger

import (
	"net/http"
	"time"

	"github.com/nuvotlyuba/Go-yandex/configs"
	"go.uber.org/zap"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		responseData := &responseData{
			status: 0,
			size:   0,
		}
		lw := loggingResponseWriter{
			ResponseWriter: w,
			responseData:   responseData,
		}

		Debug("httpRequest",
			zap.String("url", configs.ServerAddress),
			zap.String("path", r.URL.Path),
		)

		next.ServeHTTP(&lw, r)

		duration := time.Since(start)

		Debug("httpResponse",
			zap.Int("status", responseData.status),
			zap.Int("bytes", responseData.size),
			zap.Duration("elapsed", duration),
		)
	})
}
