package config

// DbConfig is a struct to store database configuration.
//
// ...
//
// Example:
//
//	...
type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	SslMode  string `mapstructure:"sslmode"`
}
