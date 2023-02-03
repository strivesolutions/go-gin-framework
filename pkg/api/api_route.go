package api

import "github.com/gin-gonic/gin"

type ApiRoute struct {
	MethodType MethodType
	Anonymous  bool
	// Deprecated: Use SkipPlanCheck instead
	SkipTrustFundCheck bool
	SkipPlanCheck      bool
	Path               string
	Handler            gin.HandlerFunc
}

func (a ApiRoute) ShouldCheckPlanId() bool {
	return !a.SkipPlanCheck && !a.SkipTrustFundCheck
}
