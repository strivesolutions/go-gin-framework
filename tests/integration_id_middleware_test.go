package tests

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/strivesolutions/go-gin-framework/pkg/api"
	"github.com/strivesolutions/go-gin-framework/pkg/server"
)

func TestIntegrationIdSuppliedGives200(t *testing.T) {
	s := server.CreateServer(server.Options{
		NoPlanIdMiddleware:        true,
		NoIntegrationIdMiddleware: false,
		HealthChecks:              passingHealthChecks(),
	})

	s.AddRoute(api.ApiRoute{
		MethodType: api.GET,
		Anonymous:  true,
		Path:       "/",
		Handler: func(ctx *gin.Context) {
			// empty response
		},
	})

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("X-Integration-Id", "1234")
	w := httptest.NewRecorder()
	s.Engine.ServeHTTP(w, req)

	ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestIntegrationIdNotSuppliedGives400(t *testing.T) {
	s := server.CreateServer(server.Options{
		NoPlanIdMiddleware:        true,
		NoIntegrationIdMiddleware: false,
		HealthChecks:              passingHealthChecks(),
	})

	s.AddRoute(api.ApiRoute{
		MethodType: api.GET,
		Anonymous:  true,
		Path:       "/",
		Handler: func(ctx *gin.Context) {
			// empty response
		},
	})

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	s.Engine.ServeHTTP(w, req)

	ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCanSkipIntegrationIdCheckOnService(t *testing.T) {
	s := server.CreateServer(server.Options{
		NoPlanIdMiddleware:        true,
		NoIntegrationIdMiddleware: true,
		HealthChecks:              passingHealthChecks(),
	})

	s.AddRoute(api.ApiRoute{
		MethodType:    api.GET,
		Anonymous:     true,
		Path:          "/",
		SkipPlanCheck: true,
		Handler: func(ctx *gin.Context) {
			// empty response
		},
	})

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	s.Engine.ServeHTTP(w, req)

	ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestIntegrationIdCanBeRead(t *testing.T) {
	s := server.CreateServer(server.Options{
		NoPlanIdMiddleware:        true,
		NoIntegrationIdMiddleware: false,
		HealthChecks:              passingHealthChecks(),
	})

	expectedIntegrationId := "1234"

	s.AddRoute(api.ApiRoute{
		MethodType: api.GET,
		Anonymous:  true,
		Path:       "/",
		Handler: func(ctx *gin.Context) {
			actualIntegrationId := api.GetIntegrationId(ctx)
			assert.Equal(t, expectedIntegrationId, actualIntegrationId)
		},
	})

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("X-Integration-Id", expectedIntegrationId)

	w := httptest.NewRecorder()
	s.Engine.ServeHTTP(w, req)

	ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestIntegrationIdCheckCanBeSkipped(t *testing.T) {
	s := server.CreateServer(server.Options{
		NoPlanIdMiddleware:        true,
		NoIntegrationIdMiddleware: false,
		HealthChecks:              passingHealthChecks(),
	})

	s.AddRoute(api.ApiRoute{
		MethodType:           api.GET,
		Anonymous:            true,
		SkipIntegrationCheck: true,
		Path:                 "/",
		Handler: func(ctx *gin.Context) {
		},
	})

	req, _ := http.NewRequest("GET", "/", nil)

	w := httptest.NewRecorder()
	s.Engine.ServeHTTP(w, req)

	ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}
