// +build integration

package rest_test

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/bygui86/go-testing/rest-examples/http-client/logging"
	"github.com/bygui86/go-testing/rest-examples/http-client/rest"
)

const (
	headerAccept          = "Accept"
	headerApplicationJson = "application/json"
)

func TestMain(m *testing.M) {
	logErr := logging.InitGlobalLogger()
	if logErr != nil {
		panic(logErr) // Panic and fail
	}

	ctx := context.Background()

	postgres, contErr := startHttpServer(ctx)
	if contErr != nil {
		panic(contErr) // Panic and fail since there is not much we can do if the container doesn't start
	}
	logging.Log.Info("HTTP server container running")
	defer stopHttpServer(postgres, ctx)

	host, port := getHostAndPort(postgres, ctx)
	logging.SugaredLog.Infof("HTTP server container exposed as: %s:%s", host, port.Port())

	setEnvVars(host, port)

	os.Exit(
		m.Run(),
	)
}

func startHttpServer(ctx context.Context) (testcontainers.Container, error) {
	logging.Log.Info("Start HTTP server")
	contReq := testcontainers.ContainerRequest{
		Image:        "bygui86/http-server:v1.0.0",
		ExposedPorts: []string{"8080/tcp"},
		Env: map[string]string{
			"JAEGER_SERVICE_NAME": "http-server",
		},
		WaitingFor: wait.ForLog("http-server up and running"),
	}

	httpServer, contErr := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: contReq,
			Started:          true,
		},
	)
	return httpServer, contErr
}

func getHostAndPort(httpServer testcontainers.Container, ctx context.Context) (string, nat.Port) {
	expPorts, expErr := httpServer.Ports(ctx)
	if expErr != nil {
		panic(expErr)
	}
	logging.Log.Debug("HTTP server exposed ports:")
	for k, v := range expPorts {
		logging.SugaredLog.Debugf("\t %s -> %v", k, v)
	}

	host, hostErr := httpServer.Host(ctx)
	if hostErr != nil {
		panic(hostErr) // Panic and fail since there is not much we can do if we cannot figure out the container ip
	}

	port, portErr := httpServer.MappedPort(ctx, "8080")
	if portErr != nil {
		panic(portErr) // Panic and fail since there is not much we can do if we cannot figure out the container port
	}
	return host, port
}

func setEnvVars(host string, port nat.Port) {
	envErr := os.Setenv("REST_SERVER_HOST", host)
	if envErr != nil {
		panic(envErr) // Panic and fail since there is not much we can do if we cannot set environment variables
	}
	envErr = os.Setenv("REST_SERVER_PORT", port.Port())
	if envErr != nil {
		panic(envErr) // Panic and fail since there is not much we can do if we cannot set environment variables
	}
}

func stopHttpServer(httpServer testcontainers.Container, ctx context.Context) {
	logging.Log.Info("HTTP server container stop")
	err := httpServer.Terminate(ctx)
	if err != nil {
		logging.SugaredLog.Error("HTTP server container stop failed: %s", err.Error())
	}
}

func TestGetProducts_Integr(t *testing.T) {
	cfg := rest.LoadConfig()

	serverUrl, serverUrlErr := rest.CreateBaseUrl(cfg.RestServerHost, cfg.RestServerPort)
	require.NoError(t, serverUrlErr)

	serverClient := rest.CreateRestClient()

	server, newErr := rest.New(cfg, serverUrl, serverClient)
	require.NoError(t, newErr)

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
