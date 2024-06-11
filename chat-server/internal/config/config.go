package config

// GRPCConfig is an interface for GRPC configuration
type GRPCConfig interface {
	Address() string
}

// PGConfig is an interface for PG configuration
type PGConfig interface {
	DSN() string
}
