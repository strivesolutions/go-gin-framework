package health

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func withTimeout(check HealthChecker, out chan HealthCheckResult) {
	defer close(out)

	r := make(chan HealthCheckResult)
	go check.Run(r)

	select {
	case <-time.After(time.Duration(check.TimeoutSeconds()) * time.Second):
		out <- Unhealthy(check.Name(), fmt.Sprintf("did not respond after %d seconds", check.TimeoutSeconds()))
	case checkResult := <-r:
		out <- checkResult
	}

}

func runChecks() ServiceHealth {
	result := CreateResponse(config.ServiceName)
	checkResults := make([]chan HealthCheckResult, len(config.Checks))

	for i, check := range config.Checks {
		checkResults[i] = make(chan HealthCheckResult)
		if check.TimeoutSeconds() > 0 {
			go withTimeout(check, checkResults[i])
		} else {
			go check.Run(checkResults[i])
		}
	}

	for i := 0; i < len(config.Checks); i++ {
		r := <-checkResults[i]
		result.AddResult(r)
		checkResults[i] = nil
	}

	return result
}

func HandleHealthRequest(ctx *gin.Context) {
	result := runChecks()

	var status int
	if result.Unhealthy {
		status = http.StatusInternalServerError
	} else {
		status = http.StatusOK
	}

	ctx.AbortWithStatusJSON(status, result)
}
