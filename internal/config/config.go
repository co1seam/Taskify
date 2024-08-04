package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config interface {
	LoadConfig() error
	LoadDatabaseConfig() error
	Get(key string) string
}

type BaseConfig struct {
	params map[string]string
}

func NewConfig() Config {
	return &BaseConfig{
		params: make(map[string]string),
	}
}

func (cfg *BaseConfig) LoadConfig() error {
	if err := InitConfig(); err != nil {
		return fmt.Errorf("error initializing configs: %v", err.Error())
	}

	cfg.params["app.port"] = viper.GetString("port")
	
	return nil
}

func (c *BaseConfig) Get(key string) string{
	return c.params[key]
}

func InitConfig() error {
	viper.AddConfigPath("golang")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %s", err)
	}

	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env variables: %s", err)
	}

	return nil;
}