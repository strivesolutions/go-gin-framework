package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHealthRequest(ctx *gin.Context, runChecks HealthChecksFunc) {
	result := runChecks()

	var status int
	if len(result.Errors) > 0 {
		status = http.StatusInternalServerError
	} else {
		status = http.StatusOK
	}

	ctx.AbortWithStatusJSON(status, result)
}
