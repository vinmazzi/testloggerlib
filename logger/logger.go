package testloggerlib

import (
	"log"
)

type LoggerObj struct {
	Status   int
	Endpoint string
	Method   string
	URL      string
}

func LoggerNew(lo LoggerObj) {
	log.Printf("%s %d %s %s", lo.Method, lo.Status, lo.Endpoint, lo.URL)
}
