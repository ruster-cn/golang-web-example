package config

import "time"

type HTTPServerConfiguration struct {
	GracefulTimeOut time.Duration `yaml:"graceful_time_out"`
	Addr            string        `yaml:"addr"`
	Port            string        `yaml:"port"`
}

func (config *HTTPServerConfiguration) Default() *HTTPServerConfiguration {
	return config
}

func (config *HTTPServerConfiguration) Validate() error {
	return nil
}
