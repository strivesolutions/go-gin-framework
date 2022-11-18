package health

import (
	"fmt"
	"net/http"
	"net/url"
)

type daprCheck struct {
	endpoint *url.URL
}

func CreateDaprHealthCheck(endpoint *url.URL) HealthChecker {
	return &daprCheck{endpoint: endpoint}
}

func (c *daprCheck) Name() string {
	return "lock-store"
}

func (c *daprCheck) TimeoutSeconds() int {
	return 2
}

func (c *daprCheck) Run(out chan HealthCheckResult) {
	defer close(out)
	const checkName = "dapr"

	url := fmt.Sprintf("%s/v1.0/healthz", c.endpoint)

	resp, err := http.Get(url)

	if err != nil || resp.StatusCode != http.StatusNoContent {
		out <- Unhealthy(checkName, fmt.Sprintf("Response from Dapr was %d", resp.StatusCode))
	} else {
		out <- Ok(checkName)
	}
}
