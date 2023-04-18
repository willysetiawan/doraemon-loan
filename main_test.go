package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"process-loan/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

var routers = config.Routers

func TestRootIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Println(err)
	} else {
		routers.ServeHTTP(rec, req)
		code := http.StatusOK
		assert.Equal(t, code, rec.Code)
		assert.Equal(t, http.StatusText(code), rec.Body.String())
	}
}
