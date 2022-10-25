package tests

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/strivesolutions/go-gin-framework/pkg/health"
	"github.com/strivesolutions/go-gin-framework/pkg/server"
)

func passingChecks() health.HealthCheckResult {
	result := health.HealthCheckResult{}
	result.AddOkResult("mock check")
	return result
}

func failingChecks() health.HealthCheckResult {
	result := health.HealthCheckResult{}
	result.AddFailedResult("mock check", "mock failure")
	return result
}

func TestHealthzHandlerPassGives200(t *testing.T) {
	s := server.CreateServer(server.Options{
		NoTrustFundMiddleware: true,
		HealthChecks:          passingChecks,
	})

	req, _ := http.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()
	s.Engine.ServeHTTP(w, req)

	ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHealthzHandlerFailGives500(t *testing.T) {
	s := server.CreateServer(server.Options{
		NoTrustFundMiddleware: true,
		HealthChecks:          failingChecks,
	})

	req, _ := http.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()
	s.Engine.ServeHTTP(w, req)

	ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
