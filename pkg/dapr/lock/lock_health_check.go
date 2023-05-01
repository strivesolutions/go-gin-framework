package lock

import (
	"fmt"

	"github.com/strivesolutions/go-gin-framework/pkg/health"
)

func lockHealthCheck(name string, out chan health.HealthCheckResult) {
	l := CreateLock("healthz")

	err := l.AcquireLock()

	if err != nil {
		out <- health.Unhealthy(name, fmt.Sprint(err))
	}
	l.Unlock()
	out <- health.Ok(name)
}

func CreateLockHealthCheck() health.HealthChecker {
	return health.CreateHealthCheckWithTimeout("lock-store", 2, lockHealthCheck)
}
