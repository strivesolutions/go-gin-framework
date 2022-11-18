package health

// To create a check without an asyncronous timeout, return 0 for TimeoutSeconds()
type HealthCheck interface {
	Name() string
	TimeoutSeconds() int
	Run(c chan HealthCheckResult)
}
