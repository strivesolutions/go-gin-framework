package state

import (
	"context"
	"fmt"

	"github.com/strivesolutions/go-gin-framework/pkg/health"
)

func storeHealthCheck(name string, out chan health.HealthCheckResult) {
	_, err := client.GetState(context.Background(), stateStoreName, "healthz", nil)

	if err != nil {
		out <- health.Unhealthy(name, fmt.Sprint(err))
	}

	out <- health.Ok(name)
}

func CreateStoreHealthCheck() health.HealthChecker {
	return health.CreateHealthCheckWithTimeout("state-store", 2, storeHealthCheck)
}
