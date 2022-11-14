package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/strivesolutions/go-gin-framework/pkg/dapr/subscribe"
)

var Subscriptions = make([]subscribe.Subscription, 0)

func HandleSubscribeRequest(ctx *gin.Context) {
	if len(Subscriptions) == 0 {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, Subscriptions)
}

func AddSubscription(sub subscribe.Subscription) {
	Subscriptions = append(Subscriptions, sub)
}
