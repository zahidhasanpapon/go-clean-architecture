package config

import "fmt"

// GetDatabaseDSN returns the database connection string
func (c *Config) GetDatabaseDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=myapp sslmode=disable",
		c.DatabaseURL, c.DatabaseURL)
}
