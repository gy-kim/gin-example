package config

import "errors"

// AppCfg is the application config
var AppCfg *Config

// Config defines the application configuration
type Config struct {
	// Address is the IP and port
	Address string

	// DSN is the data source name
	DSN string
}

// BindAddress returns the host and the port for this service
func (c *Config) BindAddress() string {
	return c.Address
}

// DataDSN returns the DSN
func (c *Config) DataDSN() string {
	return c.DSN
}

func load() error {
	return errors.New("not implemented yet")
}
