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

func TestTrustFundIdSuppliedGives200(t *testing.T) {
	s := server.CreateServer(server.Options{
		NoTrustFundMiddleware: false,
		HealthChecks:          passingHealthChecks,
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
	req.Header.Add("X-Trust-Fund-Id", "1234")
	w := httptest.NewRecorder()
	s.Engine.ServeHTTP(w, req)

	ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTrustFundIdNotSuppliedGives400(t *testing.T) {
	s := server.CreateServer(server.Options{
		NoTrustFundMiddleware: false,
		HealthChecks:          passingHealthChecks,
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
