package logger

import (
	"net/http"

	"github.com/nuvotlyuba/Go-yandex/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

func Info(msg string, field ...zapcore.Field)   {Log.Info(msg, field...)}

func Debug(msg string, fields ...zapcore.Field) {Log.Debug(msg, fields...)}

func Fatal(msg string, fields ...zapcore.Field) {Log.Fatal(msg, fields...)}


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


