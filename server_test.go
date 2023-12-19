package belajar_web

import (
	"net/http"
	"testing"
)

func Test(t *testing.T) {
	server := http.Server{
		Addr:  "localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}