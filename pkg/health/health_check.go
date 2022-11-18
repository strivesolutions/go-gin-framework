package health

// To create a check without an asyncronous timeout, return 0 for TimeoutSeconds()
// To create a simple check, you can call CreateHealthCheck or CreateHealthCheckWithTimeout
// For more control, create your own implementation of the HealthChecker interface
type HealthChecker interface {
	Name() string
	TimeoutSeconds() int
	Run(c chan HealthCheckResult)
}

type healthCheck struct {
	name           string
	timeoutSeconds int
	run            HealthCheckFunc
}

type HealthCheckFunc func(name string, out chan HealthCheckResult)

func CreateHealthCheck(name string, run HealthCheckFunc) HealthChecker {
	return &healthCheck{
		name:           name,
		timeoutSeconds: 0,
		run:            run,
	}
}

func CreateHealthCheckWithTimeout(name string, timeoutSeconds int, run HealthCheckFunc) HealthChecker {
	return &healthCheck{
		name:           name,
		timeoutSeconds: timeoutSeconds,
		run:            run,
	}
}

func (c *healthCheck) Name() string {
	return c.name
}

func (c *healthCheck) TimeoutSeconds() int {
	return c.timeoutSeconds
}

func (c *healthCheck) Run(out chan HealthCheckResult) {
	c.run(c.Name(), out)
}
