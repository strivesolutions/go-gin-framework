package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/strivesolutions/go-gin-framework/pkg/health"
)

func doNothingRouteHandler(ctx *gin.Context) {}

func passingHealthChecks() health.Config {
	return health.Config{
		ServiceName: "mock service",
	}
}

type failingCheck struct{}

func (c *failingCheck) Name() string {
	return "lock-store"
}

func (c *failingCheck) TimeoutSeconds() int {
	return 2
}

func (c *failingCheck) Run(out chan health.HealthCheckResult) {
	out <- health.Unhealthy("failed check", "mock failing check")
}

func failingHealthChecks() health.Config {
	return health.Config{
		ServiceName: "mock service",
		Checks: []health.HealthCheck{
			&failingCheck{},
		},
	}
}
