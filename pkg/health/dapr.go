package health

import (
	"fmt"
	"net/http"
)

func (result HealthCheckResult) CheckDapr(daprEndpoint string) {
	url := fmt.Sprintf("%s/v1.0/healthz", daprEndpoint)

	resp, err := http.Get(url)

	const checkName = "dapr"
	if err != nil || resp.StatusCode != http.StatusNoContent {
		result.AddFailedResult(checkName, resp.Status)
	} else {
		result.AddOkResult(checkName)
	}
}
