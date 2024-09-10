package config

// GRPCConfig is an interface for GRPC configuration
type GRPCConfig interface {
	Address() string
}

// PGConfig is an interface for PG configuration
type PGConfig interface {
	DSN() string
}

// HTTPConfig is an interface for HTTP configuration
type HTTPConfig interface {
	Address() string
}

// SwaggerConfig is an interface for HTTP configuration
type SwaggerConfig interface {
	Address() string
}

// PrometheusConfig is an interface for Prometheus configuration
type PrometheusConfig interface {
	Address() string
}
