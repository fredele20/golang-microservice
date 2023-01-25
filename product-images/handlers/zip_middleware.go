package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"
)


type GzipHandler struct {

}


func (g *GzipHandler) GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip" ){
			// create a gzipped response
			wrappedResponseWriter := NewWrappedResponseWriter(w)
			wrappedResponseWriter.Header().Set("Content-Encoding", "gzip")

			next.ServeHTTP(wrappedResponseWriter, r)
			defer wrappedResponseWriter.Flush()
			
			return
		}

		// handle normal
		next.ServeHTTP(w, r)
	})
}

type WrappedResponseWriter struct {
	w http.ResponseWriter
	gw *gzip.Writer
}

func NewWrappedResponseWriter(w http.ResponseWriter) *WrappedResponseWriter {
	gw := gzip.NewWriter(w)

	return &WrappedResponseWriter{w: w, gw: gw}
}

func (wr *WrappedResponseWriter) Header() http.Header {
	return wr.w.Header()
}

func (wr *WrappedResponseWriter) Write(d []byte) (int, error) {
	return wr.gw.Write(d)
}

func (wr *WrappedResponseWriter) WriteHeader(statuscode int) {
	wr.w.WriteHeader(statuscode)
}

func (wr *WrappedResponseWriter) Flush() {
	wr.gw.Flush()
	wr.gw.Close()
}