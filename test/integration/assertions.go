package integration

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// HasBody asserts that the request received the specified body
func HasBody(t *testing.T, req Request, body string) {
	assert.Equal(t, body, req.Body, "Should match body")
}

// HasHeader asserts that the rquest received the specified header with value
func HasHeader(t *testing.T, req Request, name string, value string) {
	checkMapOfArrayOfStrings(t, req.Headers, name, value, "header")
}

// HasMethod asserts that the request received the specified method
func HasMethod(t *testing.T, req Request, method string) {
	assert.Equal(t, method, req.Method, fmt.Sprintf("Method should be '%s'", method))
}

// HasPath asserts that the request received the specified path
func HasPath(t *testing.T, req Request, expectedPath string) {
	assert.Equal(t, expectedPath, req.Path, "Should match path")
}

// HasQueryParam checks if a request has the query parameter with the specified value
func HasQueryParam(t *testing.T, req Request, name string, value string) {
	checkMapOfArrayOfStrings(t, req.Query, name, value, "query param")
}

func checkMapOfArrayOfStrings(t *testing.T, toCheck map[string][]string, name string, value string, alias string) {
	values := toCheck[name]
	assert.NotEmpty(t, values, fmt.Sprintf("Expected %s: '%s'", alias, name))
	if len(values) > 0 {
		assert.Contains(t, values, value, fmt.Sprintf("%s '%s' should include value '%s", alias, name, value))
	}
}
