package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/strivesolutions/go-gin-framework/pkg/api"
	"github.com/strivesolutions/logger-go/pkg/logging"
)

// DetectLanguage middleware sets the locale code in the request context
// from the X-Accept-Language header.
// Making it easier to share the locale code across the application.

// supported locales
var supportedLocales = map[string]bool{
	"en": true,
	"fr": true,
}

func DetectLanguage(ctx *gin.Context) {
	value := ctx.GetHeader("X-Accept-Language")

	if value == "" {
		value = "en"
	}

	if _, ok := supportedLocales[value]; !ok {
		logging.Info("Unsupported laguage code: " + value)
		value = "en"
	}

	api.SetLocaleCode(ctx, value)
	ctx.Next()
}
