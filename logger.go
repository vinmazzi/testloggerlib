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

		log.Println("This is the STATUS CODE BEFORE THE REQUEST:", lo.status)
		n.ServeHTTP(lo, r)
		log.Println("This is the STATUS CODE AFTER THE REQUEST:", lo.status)

		wr.WriteHeader(lo.status)
		log.Printf("%s %d %s", r.Method, lo.status, r.URL.Path)
	})
}

//func LoggerNew(lo LoggerObj) {
//	log.Printf("%s %d %s %s", lo.Method, lo.Status, lo.Endpoint, lo.URL)
//}
