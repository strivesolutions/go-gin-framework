package api

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	claimsKey        = "claims"
	trustFundKey     = "trustFundId"
	planIdKey        = "planId"
	integrationIdKey = "integrationId"
)

func SetClaims(ctx *gin.Context, claims jwt.MapClaims) {
	ctx.Set(claimsKey, claims)
}

func GetClaims(ctx *gin.Context) jwt.MapClaims {
	claims, exists := ctx.Get(claimsKey)

	if exists {
		return claims.(jwt.MapClaims)
	} else {
		return jwt.MapClaims{}
	}
}

// Deprecated: Use SetPlanId instead
func SetTrustFundId(ctx *gin.Context, trustFundId int) {
	ctx.Set(trustFundKey, trustFundId)
}

// Deprecated: Use GetPlanId instead
func GetTrustFundId(ctx *gin.Context) int {
	trustFundId, exists := ctx.Get(trustFundKey)

	if !exists {
		trustFundId = 0
	}

	return trustFundId.(int)
}

func SetPlanId(ctx *gin.Context, planId string) {
	ctx.Set(planIdKey, planId)
}

func GetPlanId(ctx *gin.Context) string {
	planId, exists := ctx.Get(planIdKey)

	if !exists {
		planId = ""
	}

	return planId.(string)
}

func SetIntegrationId(ctx *gin.Context, integrationId string) {
	ctx.Set(integrationIdKey, integrationId)
}

func GetIntegrationId(ctx *gin.Context) string {
	integrationId, exists := ctx.Get(integrationIdKey)

	if !exists {
		integrationId = ""
	}

	return integrationId.(string)
}

func GetOrigin(ctx *gin.Context) string {
	return ctx.GetHeader("Origin")
}

func BearerToken(ctx *gin.Context) (string, error) {
	reqToken := ctx.GetHeader("Authorization")
	tokenSlice := strings.Split(reqToken, " ")

	if len(tokenSlice) < 2 {
		return "", errors.New("invalid authorization header format")
	} else {
		return tokenSlice[1], nil
	}
}

func UserId(ctx *gin.Context) string {
	id, exists := GetClaims(ctx)["sub"]

	if !exists {
		return ""
	}

	return id.(string)
}

func Username(ctx *gin.Context) string {
	userName, exists := GetClaims(ctx)["preferred_username"]

	if !exists {
		return ""
	}

	return userName.(string)
}

func Realm(ctx *gin.Context) string {
	userName, exists := GetClaims(ctx)["iss"]

	if !exists {
		return ""
	}

	return userName.(string)
}
