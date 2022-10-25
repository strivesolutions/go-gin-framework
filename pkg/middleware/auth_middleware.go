package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/strivesolutions/go-gin-framework/pkg/api"
	"github.com/strivesolutions/logger-go/pkg/logging"
)

type AccessResult struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

func Auth(ctx *gin.Context) {
	reqToken := ctx.GetHeader("Authorization")
	tokenSlice := strings.Split(reqToken, " ")

	if len(tokenSlice) < 2 {
		logging.Warn("Token missing or invalid for secure route")
		api.AbortUnauthorized(ctx)
		return
	}

	reqToken = tokenSlice[1]

	parser := jwt.Parser{}

	// Note:
	// This does not verify the token using the signing signature.
	// This is safe as long as this request is forwarded from a gateway which handles the actual verificiation
	token, _, err := parser.ParseUnverified(reqToken, jwt.MapClaims{})

	if err != nil {
		logging.ErrorObject(err)
		api.AbortUnauthorized(ctx)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		api.SetClaims(ctx, claims)
	}
}
