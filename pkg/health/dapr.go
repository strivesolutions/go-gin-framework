package health

import (
	"fmt"
	"net/http"
	"net/url"
)

func (result HealthCheckResult) CheckDapr(daprEndpoint *url.URL) {
	url := fmt.Sprintf("%s/v1.0/healthz", daprEndpoint)

	resp, err := http.Get(url)

	const checkName = "dapr"
	if err != nil || resp.StatusCode != http.StatusNoContent {
		result.AddFailedResult(checkName, resp.Status)
	} else {
		result.AddOkResult(checkName)
	}
}
