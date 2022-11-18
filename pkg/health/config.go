package health

import "errors"

var config Config

type Config struct {
	ServiceName string
	Checks      []HealthChecker
}

func SetConfig(c Config) {
	config = c
}

func (c Config) Validate() error {
	if c.ServiceName == "" {
		return errors.New("service name is required")
	}
	return nil
}
