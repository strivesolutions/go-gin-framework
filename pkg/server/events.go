package server

import (
	"fmt"

	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/gin-gonic/gin"
	"github.com/strivesolutions/go-gin-framework/pkg/api"
	"github.com/strivesolutions/go-gin-framework/pkg/middleware"
	"github.com/strivesolutions/go-gin-framework/pkg/pubsub"
	"github.com/strivesolutions/logger-go/pkg/logging"
)

func (s *Server) addDaprSubscribeHandler(pubsubName string) {
	s.pubsubConfigured = pubsubName != ""

	if s.pubsubConfigured {
		pubsub.PubsubName = pubsubName
		s.Engine.GET("/dapr/subscribe", middleware.HandleSubscribeRequest)
	}
}

func unwrapEvent(c *gin.Context, alwaysAck bool, handler api.EventHandlerFunc) {
	var e event.Event
	err := c.ShouldBindJSON(&e)

	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	err = handler(e)

	if err != nil {
		if alwaysAck {
			logging.Warn(fmt.Sprintf("Event handler returned error, but AlwaysAck is enabled. Message will be discarded.\n%v", err))
			c.AbortWithStatus(200)
		} else {
			c.AbortWithStatus(500)
		}
		return
	} else {
		c.AbortWithStatus(200)
	}
}

func (s *Server) AddSubscriptions(routes []api.EventRoute) {
	for _, route := range routes {
		s.AddSubscription(route)
	}
}

func (s *Server) AddSubscription(route api.EventRoute) {
	if !s.pubsubConfigured {
		logging.Error("Pubsub name was not supplied when initializing server, subscription routes cannot be added.")
		return
	}

	paths := make(map[string]bool, 0)

	if route.Subscription.Routes.Default != "" {
		paths[route.Subscription.Routes.Default] = true
	}

	for _, r := range route.Subscription.Routes.Rules {
		paths[r.Path] = true
	}

	if len(paths) == 0 {
		logging.Warn("No paths are configured for event route, events will be ignored")
		return
	}

	for p := range paths {
		s.Engine.POST(p, func(ctx *gin.Context) {
			unwrapEvent(ctx, route.AlwaysAck, route.Handler)
		})
	}

	middleware.AddSubscription(route.Subscription)
}
