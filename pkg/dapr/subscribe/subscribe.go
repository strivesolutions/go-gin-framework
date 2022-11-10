package subscribe

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetSubscriptions func() []Subscription

type Subscription struct {
	PubsubName string            `json:"pubsubname"`
	Topic      string            `json:"topic"`
	Metadata   map[string]string `json:"metadata,omitempty"`
	Routes     Routes            `json:"routes"`
}

type Routes struct {
	Rules   []Rule `json:"rules,omitempty"`
	Default string `json:"default,omitempty"`
}

type Rule struct {
	Match string `json:"match"`
	Path  string `json:"path"`
}

func HandleSubscribeRequest(ctx *gin.Context, get GetSubscriptions) {
	if get == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	result := get()

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}
