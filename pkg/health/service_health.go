package health

type ServiceHealth struct {
	ServiceName string                       `json:"serviceName"`
	Checks      map[string]HealthCheckResult `json:"checks"`
	Unhealthy   bool                         `json:"-"`
}

func CreateResponse(serviceName string) ServiceHealth {
	return ServiceHealth{
		ServiceName: serviceName,
		Checks:      map[string]HealthCheckResult{},
	}
}

func (s *ServiceHealth) AddResult(result HealthCheckResult) {
	if s.Checks == nil {
		s.Checks = map[string]HealthCheckResult{}
	}

	s.Checks[result.CheckName] = result

	if result.Status == HealthCheckStatus_Unhealthy {
		s.Unhealthy = true
	}
}
