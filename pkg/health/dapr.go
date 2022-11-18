package health

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

func CheckDapr(daprEndpoint *url.URL, wg *sync.WaitGroup, out chan HealthCheckResult) {
	defer wg.Done()
	const checkName = "dapr"

	r := make(chan HealthCheckResult, 1)

	go func() {
		url := fmt.Sprintf("%s/v1.0/healthz", daprEndpoint)

		resp, err := http.Get(url)

		const checkName = "dapr"
		if err != nil || resp.StatusCode != http.StatusNoContent {
			r <- Unhealthy(checkName, fmt.Sprintf("Response from Dapr was %d", resp.StatusCode))
		} else {
			r <- Ok(checkName)
		}
		close(r)
	}()

	select {
	case <-time.After(5 * time.Second):
		out <- Unhealthy(checkName, "Dapr did not respond within 5 seconds")
	case <-r:
		result := <-r
		out <- result
	}
}
