// +build integration

package rest_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bygui86/go-testing/rest-examples/http-server/database"
	"github.com/bygui86/go-testing/rest-examples/http-server/logging"
	"github.com/bygui86/go-testing/rest-examples/http-server/rest"
)

func TestGetProducts_Integr(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	cfg := rest.LoadConfig()
	db := database.New()
	server := rest.New(cfg, db)
	require.NotNil(t, server)

	startErr := server.Start()
	assert.NoError(t, startErr)
	assert.True(t, server.Running())

	url, urlErr := url.Parse(fmt.Sprintf("http://%s:%d/products", cfg.RestHost, cfg.RestPort))
	require.NoError(t, urlErr)
	request, reqErr := http.NewRequest(http.MethodGet, url.String(), nil)
	require.NoError(t, reqErr)
	request.Header.Set(headerAccept, headerApplicationJson)

	restClient := &http.Client{
		Timeout: 3 * time.Second,
	}
	response, respErr := restClient.Do(request)
	assert.NoError(t, respErr)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	server.Shutdown(1)
}

// TODO getProduct

// TODO createProduct

// TODO updateProduct

// TODO deleteProduct
