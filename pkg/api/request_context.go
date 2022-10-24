package api

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	claimsKey    = "claims"
	trustFundKey = "trustFundId"
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

func SetTrustFundId(ctx *gin.Context, trustFundId int) {
	ctx.Set(trustFundKey, trustFundId)
}

func GetTrustFundId(ctx *gin.Context) int {
	trustFundId, exists := ctx.Get(trustFundKey)

	if !exists {
		trustFundId = 0
	}

	return trustFundId.(int)
}

func GetOrigin(ctx *gin.Context) string {
	return ctx.GetHeader("Origin")
}
