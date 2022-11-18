package health

type HealthCheckStatus string

const (
	HealthCheckStatus_Ok        HealthCheckStatus = "ok"
	HealthCheckStatus_Unhealthy HealthCheckStatus = "unhealthy"
)
