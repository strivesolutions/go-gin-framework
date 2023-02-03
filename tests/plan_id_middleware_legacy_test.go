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

// This file tests the phase 1 backward compatibility plan w/ numeric trust fund ids

func TestCanDisablePlanIdMiddlewareWithTrustFundOption(t *testing.T) {
	s := server.CreateServer(server.Options{
		NoTrustFundMiddleware: true,
		HealthChecks:          passingHealthChecks(),
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
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestNumericPlanIdIdSetsTrustId(t *testing.T) {
	s := server.CreateServer(server.Options{
		NoPlanIdMiddleware: false,
		HealthChecks:       passingHealthChecks(),
	})

	s.AddRoute(api.ApiRoute{
		MethodType: api.GET,
		Anonymous:  true,
		Path:       "/",
		Handler: func(ctx *gin.Context) {
			trustId := api.GetTrustFundId(ctx)
			assert.NotEmpty(t, trustId)
		},
	})

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("X-Plan-Id", "1234")
	w := httptest.NewRecorder()
	s.Engine.ServeHTTP(w, req)

	ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTrustFundIdSetsPlanId(t *testing.T) {
	s := server.CreateServer(server.Options{
		NoPlanIdMiddleware: false,
		HealthChecks:       passingHealthChecks(),
	})

	s.AddRoute(api.ApiRoute{
		MethodType: api.GET,
		Anonymous:  true,
		Path:       "/",
		Handler: func(ctx *gin.Context) {
			planId := api.GetPlanId(ctx)
			assert.NotEmpty(t, planId)
		},
	})

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("X-Trust-Fund-Id", "1234")
	w := httptest.NewRecorder()
	s.Engine.ServeHTTP(w, req)

	ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}
