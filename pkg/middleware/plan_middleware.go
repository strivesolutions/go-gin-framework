package middleware

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/strivesolutions/go-gin-framework/pkg/api"
	"github.com/strivesolutions/logger-go/pkg/logging"
)

// Migration path for trust fund id deprecation:
// Phase 1 - Both middlewares are included in pipeline. Trust ID header sets Plan ID.  Plan ID will set Trust ID header, IF plan ID is numeric.
//			 At least one of the headers is required.
// Phase 2 - Trust fund references disappear and plan ID header is required.

func Plan(ctx *gin.Context) {
	getPlanId(ctx)
	getTrustFundId(ctx)
	getPlanName(ctx)

	planId := api.GetPlanId(ctx)
	trustFundId := api.GetTrustFundId(ctx)

	if planId == "" && trustFundId == 0 {
		err := errors.New("header x-trust-fund-id or x-plan-id is required")
		logging.ErrorObject(err)
		api.AbortBadRequest(ctx, err)
		return
	}

	// plan id is set, trust fund is not
	if planId != "" && trustFundId == 0 {
		// compatiblity - set trust fund if planId is numeric
		trustFundId, err := strconv.Atoi(planId)
		if err == nil {
			api.SetTrustFundId(ctx, trustFundId)
		}
	}

	// trust fund is set, plan id is not
	if trustFundId != 0 && planId == "" {
		api.SetPlanId(ctx, strconv.Itoa(trustFundId))
	}
}

func getPlanId(ctx *gin.Context) {
	value := ctx.GetHeader("X-Plan-Id")

	api.SetPlanId(ctx, value)
}

func getPlanName(ctx *gin.Context) {
	value := ctx.GetHeader("X-Plan-Name")

	api.SetPlanName(ctx, value)

}

// Deprecated: to be replaced by PlanId
func getTrustFundId(ctx *gin.Context) {

	value := ctx.GetHeader("X-Trust-Fund-Id")

	trustFundId, err := strconv.Atoi(value)

	if err != nil {
		return
	}

	api.SetTrustFundId(ctx, trustFundId)
}
