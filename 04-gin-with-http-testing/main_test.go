package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gin-http-server/router"

	"github.com/gavv/httpexpect"
)

func TestRouter(t *testing.T) {
	// create http.Handler
	handler := router.Init()

	// run server using httptest
	server := httptest.NewServer(handler)
	defer server.Close()

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)

	// is it working?
	e.GET("/ping").
		Expect().
		Status(http.StatusOK)
}
