package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndpoints(t *testing.T) {

	endpoints := []struct {
		route              string
		method             string
		expectedStatusCode int
	}{
		{
			route:              "/",
			method:             "GET",
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, endpoint := range endpoints {
		req, _ := http.NewRequest(
			endpoint.method,
			endpoint.route,
			nil,
		)

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
