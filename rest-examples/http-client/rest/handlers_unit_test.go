// +build unit !integration

package rest_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bygui86/go-testing/rest-examples/http-client/commons"
	"github.com/bygui86/go-testing/rest-examples/http-client/logging"
	"github.com/bygui86/go-testing/rest-examples/http-client/rest"
)

// MockClient is the mocked HTTPClient
type MockClient struct {
	responseMocker func(req *http.Request) (*http.Response, error)
}

// Do is the mocked HTTPClient's `Do` func
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.responseMocker(req)
}

func TestGetProducts_Unit_Success_Empty(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	cfg := rest.LoadConfig()

	baseUrl, urlErr := rest.CreateBaseUrl(cfg.RestServerHost, cfg.RestServerPort)
	require.NoError(t, urlErr)

	// response JSON
	responseJson := `[]`
	// create a new reader with JSON response
	readCloser := ioutil.NopCloser(
		bytes.NewReader([]byte(responseJson)),
	)
	client := &MockClient{
		responseMocker: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       readCloser,
			}, nil
		},
	}

	server, newErr := rest.New(cfg, baseUrl, client)
	require.NoError(t, newErr)

	request, err := http.NewRequest("GET", "/products", nil)
	assert.NoError(t, err)

	responseRec := httptest.NewRecorder()
	http.HandlerFunc(server.GetProducts).ServeHTTP(responseRec, request)
	assert.Equal(t, http.StatusOK, responseRec.Code)
}

func TestGetProducts_Unit_Success_Single(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	cfg := rest.LoadConfig()

	baseUrl, urlErr := rest.CreateBaseUrl(cfg.RestServerHost, cfg.RestServerPort)
	require.NoError(t, urlErr)

	// response JSON
	responseJson := `[{"id":1, "name":"sample", "price":9.90}]`
	// create a new reader with JSON response
	readCloser := ioutil.NopCloser(
		bytes.NewReader([]byte(responseJson)),
	)
	client := &MockClient{
		responseMocker: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       readCloser,
			}, nil
		},
	}

	server, newErr := rest.New(cfg, baseUrl, client)
	require.NoError(t, newErr)

	request, err := http.NewRequest("GET", "/products", nil)
	assert.NoError(t, err)

	responseRec := httptest.NewRecorder()
	http.HandlerFunc(server.GetProducts).ServeHTTP(responseRec, request)
	assert.Equal(t, http.StatusOK, responseRec.Code)
	var products []*commons.Product
	unmarshErr := json.NewDecoder(responseRec.Body).Decode(&products)
	assert.NoError(t, unmarshErr)
	assert.NotNil(t, products)
	assert.Len(t, products, 1)
}

func TestGetProducts_Unit_Success_Multi(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	cfg := rest.LoadConfig()

	baseUrl, urlErr := rest.CreateBaseUrl(cfg.RestServerHost, cfg.RestServerPort)
	require.NoError(t, urlErr)

	// response JSON
	responseJson := `[{"id":1, "name":"sample", "price":9.90}, {"id":2, "name":"sample-2", "price":42.90}, {"id":3, "name":"sample-3", "price":0.90}]`
	// create a new reader with JSON response
	readCloser := ioutil.NopCloser(
		bytes.NewReader([]byte(responseJson)),
	)
	client := &MockClient{
		responseMocker: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       readCloser,
			}, nil
		},
	}

	server, newErr := rest.New(cfg, baseUrl, client)
	require.NoError(t, newErr)

	request, err := http.NewRequest("GET", "/products", nil)
	assert.NoError(t, err)

	responseRec := httptest.NewRecorder()
	http.HandlerFunc(server.GetProducts).ServeHTTP(responseRec, request)
	assert.Equal(t, http.StatusOK, responseRec.Code)
	var products []*commons.Product
	unmarshErr := json.NewDecoder(responseRec.Body).Decode(&products)
	assert.NoError(t, unmarshErr)
	assert.NotNil(t, products)
	assert.Len(t, products, 3)
}

func TestGetProducts_Unit_Fail(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	cfg := rest.LoadConfig()

	baseUrl, urlErr := rest.CreateBaseUrl(cfg.RestServerHost, cfg.RestServerPort)
	require.NoError(t, urlErr)

	client := &MockClient{
		responseMocker: func(*http.Request) (*http.Response, error) {
			return nil, errors.New("error from web server")
		},
	}

	server, newErr := rest.New(cfg, baseUrl, client)
	require.NoError(t, newErr)

	request, err := http.NewRequest("GET", "/products", nil)
	assert.NoError(t, err)

	responseRec := httptest.NewRecorder()
	http.HandlerFunc(server.GetProducts).ServeHTTP(responseRec, request)
	assert.Equal(t, http.StatusInternalServerError, responseRec.Code)
}

// TODO getProduct

// TODO createProduct

// TODO updateProduct

// TODO deleteProduct
