package server

import (
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/gin-gonic/gin"
	"github.com/strivesolutions/go-gin-framework/pkg/api"
	"github.com/strivesolutions/go-gin-framework/pkg/middleware"
)

func unwrapEvent(c *gin.Context, alwaysAck bool, handler api.EventHandlerFunc) {
	var e event.Event
	err := c.ShouldBindJSON(&e)

	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	err = handler(e)

	if err != nil && !alwaysAck {
		c.AbortWithStatus(500)
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
	s.Engine.POST(route.Path)
	middleware.AddSubscription(route.Subscription)
}
