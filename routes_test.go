package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndpoints(t *testing.T) {

	quote := getStubQuote()
	quoteJson, _ := json.Marshal(quote)

	endpoints := []struct {
		route              string
		method             string
		expectedStatusCode int
		body               []byte
	}{
		{
			route:              "/",
			method:             "GET",
			expectedStatusCode: http.StatusOK,
			body:               nil,
		},
		{
			route:              "/quotes",
			method:             "POST",
			expectedStatusCode: http.StatusCreated,
			body:               quoteJson,
		},
	}

	for _, endpoint := range endpoints {
		req, _ := http.NewRequest(
			endpoint.method,
			endpoint.route,
			bytes.NewBuffer(endpoint.body),
		)

		req.Header.Add("Content-Type", "application/json;charset=utf-8")

		app := Setup()
		res, _ := app.Test(req, -1)

		assert.Equal(t, res.StatusCode, endpoint.expectedStatusCode)
	}
}

func Test404(t *testing.T) {

	req, _ := http.NewRequest(
		"GET",
		"/invalid",
		nil,
	)

	app := Setup()
	res, _ := app.Test(req, -1)

	assert.Equal(t, res.StatusCode, http.StatusNotFound)
}
