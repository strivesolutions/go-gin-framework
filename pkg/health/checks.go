package health

type HealthCheckStatus string

const (
	HealthCheckStatus_Ok   = "ok"
	HealthCheckStatus_Fail = "fail"
)

type HealthCheckResult struct {
	ServiceName string                       `json:"serviceName"`
	Checks      map[string]HealthCheckStatus `json:"checks"`
	Errors      []string                     `json:"errors"`
}

type HealthChecksFunc func() HealthCheckResult

func (h *HealthCheckResult) AddOkResult(name string) {
	if h.Checks == nil {
		h.Checks = make(map[string]HealthCheckStatus)
	}
	h.Checks[name] = HealthCheckStatus_Ok
}

func (h *HealthCheckResult) AddFailedResult(name string, errorDetail string) {
	if h.Checks == nil {
		h.Checks = make(map[string]HealthCheckStatus)
	}
	h.Checks[name] = HealthCheckStatus_Fail
	h.Errors = append(h.Errors, errorDetail)
}
