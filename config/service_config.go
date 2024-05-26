package config

// ServiceConfig is a struct to store web service configuration.
//
// Bind is the address to allow connections from.
// Port is the port to listen on.
//
// Example:
//
//	bind: "0.0.0.0"
//	port: 9000
type ServiceConfig struct {
	Bind string `mapstructure:"bind"`
	Port int    `mapstructure:"port"`
}
