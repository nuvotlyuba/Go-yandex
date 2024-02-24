package logger

import (
	"net/http"
	"time"

	"github.com/nuvotlyuba/Go-yandex/configs"
	"go.uber.org/zap"
)
type (
	responseData struct {
		status int
		size int
	}

	loggingResponseWriter struct {
		http.ResponseWriter
		responseData *responseData
	}
)

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size +=size
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
    r.ResponseWriter.WriteHeader(statusCode)
    r.responseData.status = statusCode
}

var Log *zap.Logger = zap.NewNop()


func Initialize(level string, appEnv string) error {

	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return err
	}
	var cfg zap.Config
	if configs.Stage(appEnv) == configs.Production {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}
	cfg.Level = lvl

	zl, err := cfg.Build()
	if err != nil {
		return err
	}

	Log = zl

	return nil
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func( w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		responseData := &responseData{
			status: 0,
			size:0,
		}
		lw := loggingResponseWriter {
			ResponseWriter: w,
			responseData: responseData,
		}

		Log.Debug("httpRequest",
			zap.String("url", configs.ServerAddress),
			zap.String("path", r.URL.Path),
		)

		next.ServeHTTP(&lw, r)

		duration := time.Since(start)

		Log.Debug("httpResponse",
			zap.Int("status", responseData.status),
			zap.Int("bytes", responseData.size),
			zap.Duration("elapsed", duration),
		)
	})
}

