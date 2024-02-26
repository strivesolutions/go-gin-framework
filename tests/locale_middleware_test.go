package tests

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/strivesolutions/go-gin-framework/pkg/api"
	"github.com/strivesolutions/go-gin-framework/pkg/server"
)

func TestLocaleMiddleware_WhenLocaleCodeIsSet_ShouldReturn200(t *testing.T) {
	s := server.CreateServer(server.Options{
		NoPlanIdMiddleware:        true,
		NoIntegrationIdMiddleware: true,
		HealthChecks:              passingHealthChecks(),
	})

	s.AddRoute(api.ApiRoute{
		MethodType: api.GET,
		Anonymous:  true,
		Path:       "/",
		Handler: func(ctx *gin.Context) {
			localeCode := api.GetLocaleCode(ctx)

			api.OkResponse(ctx, localeCode)
		},
	})

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("X-Accept-Language", "en")
	w := httptest.NewRecorder()
	s.Engine.ServeHTTP(w, req)

	body, _ := io.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"data":"en","error":null}`, string(body))
}

func TestLocaleMiddleware_WhenLocaleIsNotInHeader_ShouldReturn200(t *testing.T) {
	s := server.CreateServer(server.Options{
		NoPlanIdMiddleware:        true,
		NoIntegrationIdMiddleware: true,
		HealthChecks:              passingHealthChecks(),
	})

	s.AddRoute(api.ApiRoute{
		MethodType: api.GET,
		Anonymous:  true,
		Path:       "/",
		Handler: func(ctx *gin.Context) {
			localeCode := api.GetLocaleCode(ctx)

			api.OkResponse(ctx, localeCode)
		},
	})

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	s.Engine.ServeHTTP(w, req)

	body, _ := io.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"data":"en","error":null}`, string(body))
}

func TestLocaleMiddleware_WhenLocaleIsNotSupported_ShouldReturn200(t *testing.T) {
	s := server.CreateServer(server.Options{
		NoPlanIdMiddleware:        true,
		NoIntegrationIdMiddleware: true,
		HealthChecks:              passingHealthChecks(),
	})

	s.AddRoute(api.ApiRoute{
		MethodType: api.GET,
		Anonymous:  true,
		Path:       "/",
		Handler: func(ctx *gin.Context) {
			localeCode := api.GetLocaleCode(ctx)

			api.OkResponse(ctx, localeCode)
		},
	})

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("X-Accept-Language", "es")
	w := httptest.NewRecorder()
	s.Engine.ServeHTTP(w, req)

	body, _ := io.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"data":"en","error":null}`, string(body))

}
