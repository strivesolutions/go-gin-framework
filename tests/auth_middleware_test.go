package tests

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/strivesolutions/go-gin-framework/pkg/api"
	"github.com/strivesolutions/go-gin-framework/pkg/middleware"
	"github.com/strivesolutions/go-gin-framework/pkg/server"
)

func TestAuthMiddlewareCalledForAuthedRoute(t *testing.T) {
	authCalled := false
	fakeAuthMiddleware := func(ctx *gin.Context) {
		authCalled = true
	}

	server.AuthMiddleware = fakeAuthMiddleware

	s := server.CreateServer(server.Options{
		NoTrustFundMiddleware: true,
		HealthChecks:          passingHealthChecks,
	})

	s.AddRoute(api.ApiRoute{
		MethodType: api.GET,
		Path:       "/",
		Anonymous:  false,
		Handler:    doNothingRouteHandler,
	})

	req, _ := http.NewRequest("GET", "/", nil)

	w := httptest.NewRecorder()
	s.Engine.ServeHTTP(w, req)

	assert.True(t, authCalled)
}

func TestAuthMiddlewareNotCalledForAnonymousRoute(t *testing.T) {
	authCalled := false
	fakeAuthMiddleware := func(ctx *gin.Context) {
		authCalled = true
	}

	server.AuthMiddleware = fakeAuthMiddleware

	s := server.CreateServer(server.Options{
		NoTrustFundMiddleware: true,
		HealthChecks:          passingHealthChecks,
	})

	s.AddRoute(api.ApiRoute{
		MethodType: api.GET,
		Path:       "/",
		Anonymous:  true,
		Handler:    doNothingRouteHandler,
	})

	req, _ := http.NewRequest("GET", "/", nil)

	w := httptest.NewRecorder()
	s.Engine.ServeHTTP(w, req)

	assert.False(t, authCalled)
}

func TestAuthMiddlewareCalledByDefault(t *testing.T) {
	authCalled := false
	fakeAuthMiddleware := func(ctx *gin.Context) {
		authCalled = true
	}

	server.AuthMiddleware = fakeAuthMiddleware

	s := server.CreateServer(server.Options{
		NoTrustFundMiddleware: true,
		HealthChecks:          passingHealthChecks,
	})

	s.AddRoute(api.ApiRoute{
		MethodType: api.GET,
		Path:       "/",
		//Anonymous:  NOT SET,
		Handler: doNothingRouteHandler,
	})

	req, _ := http.NewRequest("GET", "/", nil)

	w := httptest.NewRecorder()
	s.Engine.ServeHTTP(w, req)

	assert.True(t, authCalled)
}

func TestAuthMiddlewareFailureHaltsHandling(t *testing.T) {
	r := gin.Default()

	responseString := "Should't see me"

	mockHandler := func(ctx *gin.Context) {
		ctx.String(200, "%s", responseString)
	}

	r.GET("/", middleware.Auth, mockHandler)

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.NotContains(t, string(responseData), responseString)
}
