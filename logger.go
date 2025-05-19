package testloggerlib

import (
	"log"
	"net/http"
)

type LoggerObj struct {
	http.ResponseWriter
	status int
}

func (lo *LoggerObj) WriteHeader(status int) {
	lo.status = status
	lo.ResponseWriter.WriteHeader(status)
}

func Middleware(n http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		lo := &LoggerObj{
			ResponseWriter: wr,
		}

		log.Printf("Type %T", wr)
		n.ServeHTTP(lo, r)
		log.Printf("%s %d %s", r.Method, lo.status, r.URL.Path)
	})
}
