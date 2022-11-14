package tests

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/stretchr/testify/assert"
	"github.com/strivesolutions/go-gin-framework/pkg/api"
	"github.com/strivesolutions/go-gin-framework/pkg/dapr/subscribe"
	"github.com/strivesolutions/go-gin-framework/pkg/server"
)

func TestSubscribeHandlerReturns404WhenNotConfigured(t *testing.T) {
	s := server.CreateServer(server.Options{
		NoTrustFundMiddleware: true,
		HealthChecks:          passingHealthChecks,
		PubsubName:            "mock-pubsub",
	})

	// no subscriptions added

	req, _ := http.NewRequest("GET", "/dapr/subscribe", nil)
	w := httptest.NewRecorder()
	s.Engine.ServeHTTP(w, req)

	ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestSubscribeHandlerReturns200WhenConfigured(t *testing.T) {
	s := server.CreateServer(server.Options{
		NoTrustFundMiddleware: true,
		HealthChecks:          passingHealthChecks,
		PubsubName:            "mock-pubsub",
	})

	s.AddSubscription(api.EventRoute{
		AlwaysAck: false,
		Handler:   func(e event.Event) *api.EventHandlerError { return nil },
		Subscription: subscribe.Subscription{
			Routes: subscribe.Routes{
				Default: "mock/route",
			},
		},
	})
	req, _ := http.NewRequest("GET", "/dapr/subscribe", nil)
	w := httptest.NewRecorder()
	s.Engine.ServeHTTP(w, req)

	ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}
