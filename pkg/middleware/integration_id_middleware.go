package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/strivesolutions/go-gin-framework/pkg/api"
	"github.com/strivesolutions/logger-go/pkg/logging"
)

func IntegrationId(ctx *gin.Context) {
	value := ctx.GetHeader("X-Integration-Id")

	api.SetIntegrationId(ctx, value)

	integrationId := api.GetIntegrationId(ctx)

	if integrationId == "" {
		err := errors.New("header x-integration-id is required")
		logging.ErrorObject(err)
		api.AbortBadRequest(ctx, err)
		return
	}
}
