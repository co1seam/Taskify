package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func (cfg *BaseConfig) LoadDatabaseConfig() error {
	if err := InitConfig(); err != nil {
		return fmt.Errorf("error initializing database config: %v", err)
	}
	  
	cfg.params["db.host"] = viper.GetString("db.host")
	cfg.params["db.port"] = viper.GetString("db.port")
	cfg.params["db.username"] = viper.GetString("db.username")
	cfg.params["db.dbname"] = viper.GetString("db.dbname")
	cfg.params["db.sslmode"] = viper.GetString("db.sslmode")
	  
	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return fmt.Errorf("DB_PASSWORD environment variable not set")
	}
	cfg.params["db.password"] = dbPassword
	
	return nil
}