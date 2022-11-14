package api

import (
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/strivesolutions/go-gin-framework/pkg/dapr/subscribe"
)

type EventHandlerFunc func(e event.Event) error

type EventRoute struct {
	AlwaysAck    bool
	Path         string
	Handler      EventHandlerFunc
	Subscription subscribe.Subscription
}

// func CreateEventRoute(path string, alwaysAck bool, handler EventHandlerFunc) EventRoute {
// 	return ApiRoute{
// 		MethodType:         POST,
// 		Anonymous:          true,
// 		SkipTrustFundCheck: true,
// 		Path:               path,
// 		Handler: func(c *gin.Context) {
// 			unwrapEvent(c, alwaysAck, handler)
// 		},
// 	}
// }
