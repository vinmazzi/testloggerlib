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
}

func Middleware(n http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		lo := &LoggerObj{
			ResponseWriter: wr,
		}

		n.ServeHTTP(lo, r)

		log.Println(r.Header)
		log.Printf("%s %d %s", r.Method, lo.status, r.URL.Path)
	})
}

//func LoggerNew(lo LoggerObj) {
//	log.Printf("%s %d %s %s", lo.Method, lo.Status, lo.Endpoint, lo.URL)
//}
