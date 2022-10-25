package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/strivesolutions/go-gin-framework/pkg/api"
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
