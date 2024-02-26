package middleware

import (
	"strings"

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

var defaultLocale = "en"

func DetectLanguage(ctx *gin.Context) {
	value := strings.ToLower(ctx.GetHeader("X-Accept-Language"))

	if value == "" {
		value = defaultLocale
	}

	if _, ok := supportedLocales[value]; !ok {
		logging.Info("Unsupported language code: " + value)
		value = defaultLocale
	}

	api.SetLocaleCode(ctx, value)
	ctx.Next()
}
