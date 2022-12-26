package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"
)

// GZipHandler  is a handler for compressing files
type GZipHandler struct{}

// GzipMiddleware is a middleware for compressing files
func (g *GZipHandler) GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			// create a gzipped response
			wrw := NewResponseWriteWrapper(rw)
			// set header
			wrw.Header().Set("Content-Encoding", "gzip")
			next.ServeHTTP(wrw, r)
			defer wrw.Flush()

			return
		}

		// handle normal
		next.ServeHTTP(rw, r)
	})
}

// ResponseWriterWrapper wraps an http.ResponseWriter
type ResponseWriterWrapper struct {
	rw http.ResponseWriter
	gw *gzip.Writer
}

// NewResponseWriteWrapper creates a new ResponseWriterWrapper
func NewResponseWriteWrapper(rw http.ResponseWriter) *ResponseWriterWrapper {
	gw := gzip.NewWriter(rw)

	return &ResponseWriterWrapper{
		rw: rw,
		gw: gw,
	}
}

// Header returns a ResponseWriters header
func (wr *ResponseWriterWrapper) Header() http.Header {
	return wr.rw.Header()
}

// Write returns gzip.Write
func (wr *ResponseWriterWrapper) Write(d []byte) (int, error) {
	return wr.gw.Write(d)
}

// WriteHeader sends a ResponseWriters response header with the provided status code
func (wr *ResponseWriterWrapper) WriteHeader(statusCode int) {
	wr.rw.WriteHeader(statusCode)
}

// Flush flushes data and close the gzip.Writer
func (wr *ResponseWriterWrapper) Flush() {
	wr.gw.Flush()
	wr.gw.Close()
}
