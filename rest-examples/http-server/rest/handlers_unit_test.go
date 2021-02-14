// +build unit !integration

package rest_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bygui86/go-testing/rest-examples/http-server/database"
	"github.com/bygui86/go-testing/rest-examples/http-server/logging"
	"github.com/bygui86/go-testing/rest-examples/http-server/rest"
)

func TestGetProducts_Unit(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	cfg := &rest.Config{
		RestHost: "localhost",
		RestPort: 8080,
	}
	db := database.New() // TODO replace with a mock
	server := rest.New(cfg, db)
	require.NotNil(t, server)

	request, err := http.NewRequest("GET", "/products", nil)
	require.NoError(t, err)

	responseRec := httptest.NewRecorder()
	http.HandlerFunc(server.GetProducts).ServeHTTP(responseRec, request)
	assert.Equal(t, http.StatusOK, responseRec.Code)
}

// TODO getProduct

// TODO createProduct

// TODO updateProduct

// TODO deleteProduct
