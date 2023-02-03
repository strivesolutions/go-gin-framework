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

func TestPlanIdIdSuppliedGives200(t *testing.T) {
	s := server.CreateServer(server.Options{
		NoPlanIdMiddleware: false,
		HealthChecks:       passingHealthChecks(),
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
	req.Header.Add("X-Plan-Id", "1234")
	w := httptest.NewRecorder()
	s.Engine.ServeHTTP(w, req)

	ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPlanIdNotSuppliedGives400(t *testing.T) {
	s := server.CreateServer(server.Options{
		NoPlanIdMiddleware: false,
		HealthChecks:       passingHealthChecks(),
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

func TestCanSkipPlanIdCheckOnRoute(t *testing.T) {
	s := server.CreateServer(server.Options{
		NoPlanIdMiddleware: false,
		HealthChecks:       passingHealthChecks(),
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
