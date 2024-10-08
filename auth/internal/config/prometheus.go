package config

import (
	"errors"
	"net"
	"os"
)

const (
	prometheusHostEnvName = "PROMETHEUS_HOST"
	prometheusPortEnvName = "PROMETHEUS_PORT"
)

type prometheusConfig struct {
	host string
	port string
}

// NewPrometheusConfig creates a new Prometheus configuration by reading the host and port
func NewPrometheusConfig() (PrometheusConfig, error) {
	host := os.Getenv(prometheusHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("prometheus host not found")
	}

	port := os.Getenv(prometheusPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("prometheus port not found")
	}

	return &prometheusConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *prometheusConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
