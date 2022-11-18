package health

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/strivesolutions/logger-go/pkg/logging"
)

func withTimeout(check HealthCheck, out chan HealthCheckResult) {
	defer close(out)

	r := make(chan HealthCheckResult)
	go check.Run(r)

	select {
	case <-time.After(time.Duration(check.TimeoutSeconds()) * time.Second):
		out <- Unhealthy(check.Name(), fmt.Sprintf("Check did not respond after %d seconds", check.TimeoutSeconds()))
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
		logging.Info(fmt.Sprintf("Started check %d", i))
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
