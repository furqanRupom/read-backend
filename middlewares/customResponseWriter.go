package middlewares

import (
	"bytes"
	"io"
	"net/http"
)

type CustomResponseWriter struct {
	http.ResponseWriter
	body       []byte
	statusCode int
	header     http.Header
}

func NewCustomResponseWriter(rw http.ResponseWriter) *CustomResponseWriter {
	return &CustomResponseWriter{
		ResponseWriter: rw,
		header:         rw.Header(),
		body:           make([]byte, 0),
	}
}

func (w *CustomResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *CustomResponseWriter) GetResponse() http.Response {
	return http.Response{
		Status:        http.StatusText(w.statusCode),
		StatusCode:    w.statusCode,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          io.NopCloser(bytes.NewBuffer(w.body)),
		ContentLength: int64(len(w.body)),
		Header:        w.header,
	}
}

func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	written, err := w.ResponseWriter.Write(b)
	if err == nil {
		writtenBytes := b[:written]
		w.body = append(w.body, writtenBytes...)
	}
	return written, err
}

func (w *CustomResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
