package main

import (
	"net/http"

	"github.com/Djarvur/go-mk-ms01-swagger/restapi/operations/ping"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// Responder todo
type Responder struct {
	code     int
	response interface{}
	headers  http.Header
}

// PingHandlerFunc todo
func PingHandlerFunc(ping.PingParams) middleware.Responder {
	return &Responder{
		http.StatusOK,
		"pong",
		make(http.Header),
	}
}

// WriteResponse todo
func (r *Responder) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	for k, v := range r.headers {
		for _, val := range v {
			rw.Header().Add(k, val)
		}
	}

	rw.WriteHeader(r.code)

	if r.response != nil {
		if err := producer.Produce(rw, r.response); err != nil {
			panic(err)
		}
	}
}
