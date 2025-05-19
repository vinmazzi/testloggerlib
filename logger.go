package testloggerlib

import (
	"log"
	"net/http"
)

type LoggerObj struct {
	http.ResponseWriter
	Status int
}

func (lo *LoggerObj) WriteHeader(status int) {
	lo.Status = status
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		lo := &LoggerObj{
			ResponseWriter: wr,
			Status:         http.StatusOK,
		}

		next.ServeHTTP(lo, r)

		log.Printf("%s %d %s %s", r.Method, lo.Status, r.URL.Path, r.URL)
	})
}

//func LoggerNew(lo LoggerObj) {
//	log.Printf("%s %d %s %s", lo.Method, lo.Status, lo.Endpoint, lo.URL)
//}
