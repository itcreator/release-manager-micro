package config

import "os"

//AppConfig is an interface for application configuration
type AppConfig interface {
	GetDatabaseHost() string
	GetDatabaseName() string
	GetDatabaseUser() string
	GetDatabasePassword() string
}

type appConfig struct {
	databaseHost     string
	databaseName     string
	databaseUser     string
	databasePassword string
	httpPort         string
}

//New is constructor for application config
func New() AppConfig {
	config := new(appConfig)
	config.load()

	return config
}

func (c *appConfig) load() {
	c.databaseHost = os.Getenv("APP_DATABASE_HOST")
	c.databaseName = os.Getenv("POSTGRES_DB")
	c.databaseUser = os.Getenv("POSTGRES_USER")
	c.databasePassword = os.Getenv("POSTGRES_PASSWORD")
}

func (c *appConfig) GetDatabaseHost() string {
	return c.databaseHost
}

func (c *appConfig) GetDatabaseName() string {
	return c.databaseName
}

func (c *appConfig) GetDatabaseUser() string {
	return c.databaseUser
}

func (c *appConfig) GetDatabasePassword() string {
	return c.databasePassword
}
