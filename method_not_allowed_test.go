package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMethodNotAllowed(t *testing.T) {
	router := httprouter.New()

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, R *http.Request) {
		fmt.Fprint(w, "Gak Boleh")
	})

	router.POST("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "POST")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Gak Boleh", string(body))
}
