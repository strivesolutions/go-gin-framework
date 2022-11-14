package api

import (
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/strivesolutions/go-gin-framework/pkg/dapr/subscribe"
)

type EventHandlerError struct {
	Error    error
	CanRetry bool
}

type EventHandlerFunc func(e event.Event) *EventHandlerError

type EventRoute struct {
	AlwaysAck    bool
	Handler      EventHandlerFunc
	Subscription subscribe.Subscription
}
