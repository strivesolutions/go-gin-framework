package health

type HealthCheckResult struct {
	CheckName    string            `json:"-"`
	Status       HealthCheckStatus `json:"status"`
	ErrorDetails string            `json:"errorDetails,omitempty"`
}

func Ok(checkName string) HealthCheckResult {
	return HealthCheckResult{
		CheckName: checkName,
		Status:    HealthCheckStatus_Ok,
	}
}

func Unhealthy(checkName, errorDetails string) HealthCheckResult {
	return HealthCheckResult{
		CheckName:    checkName,
		Status:       HealthCheckStatus_Unhealthy,
		ErrorDetails: errorDetails,
	}
}
