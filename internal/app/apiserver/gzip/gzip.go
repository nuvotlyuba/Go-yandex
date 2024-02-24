package gzip

import (
	"net/http"
	"strings"
)


func GzipMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		ow := w

		acceptEncoding := r.Header.Get("Accept-Encoding")
		isSupportsGzip := strings.Contains(acceptEncoding, "gzip")

		currentContentType := r.Header.Get("Content-Type")
		isSupportsContentType := strings.Contains(currentContentType, "application/json") ||
			strings.Contains(currentContentType, "text/html")

		if isSupportsGzip && isSupportsContentType {
			cw := newCompressWriter(w)
			ow = cw

			defer cw.Close()
		}

		contentEncoding := r.Header.Get("Content-Encoding")
		sendsGzip := strings.Contains(contentEncoding, "gzip")
		if sendsGzip {
			cr, err := newCompressReader(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			r.Body = cr
			defer cr.Close()
		}
		next.ServeHTTP(ow, r)
	})
}

