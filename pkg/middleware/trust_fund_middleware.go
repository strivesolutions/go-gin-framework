package middleware

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/strivesolutions/go-gin-framework/pkg/api"
	"github.com/strivesolutions/logger-go/pkg/logging"
)

func TrustFundId(ctx *gin.Context) {

	value := ctx.GetHeader("X-Trust-Fund-Id")

	trustFundId, err := strconv.Atoi(value)

	if err != nil {
		err = errors.New("header x-trust-fund-id is required")
		logging.ErrorObject(err)
		api.AbortBadRequest(ctx, err)
		return
	}

	api.SetTrustFundId(ctx, trustFundId)
}
