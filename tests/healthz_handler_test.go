package tests

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/strivesolutions/go-gin-framework/pkg/server"
)

func TestHealthzHandlerPassGives200(t *testing.T) {
	s := server.CreateServer(server.Options{
		NoTrustFundMiddleware: true,
		HealthChecks:          passingHealthChecks(),
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
		HealthChecks:          failingHealthChecks(),
	})

	req, _ := http.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()
	s.Engine.ServeHTTP(w, req)

	ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestHealthzHandlerIgnoresTrustMiddleware(t *testing.T) {
	s := server.CreateServer(server.Options{
		NoTrustFundMiddleware: false,
		HealthChecks:          passingHealthChecks(),
	})

	req, _ := http.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()
	s.Engine.ServeHTTP(w, req)

	ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}
