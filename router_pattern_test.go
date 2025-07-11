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

func TestRouterPatternNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/item/:itemId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		text := "Product " + p.ByName("id") + " Item " + p.ByName("itemId")
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/products/1/item/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 1 Item 1", string(body))
}

func TestRouterPatternCatchAllParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		text := "Image : " + p.ByName("image")
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Image : /small/profile.png", string(body))
}
