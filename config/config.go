package config

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	ENVIRONMENT string `mapstructure:"environment"`
	PORT        string `mapstructure:"port"`

	DBUSER           string `mapstructure:"pg_user"`
	DBPASS           string `mapstructure:"pg_pw"`
	DBHOST           string `mapstructure:"pg_host"`
	DBNAME           string `mapstructure:"pg_db"`
	MAXDBCONNECTIONS string `mapstructure:"max_db_conns"`

	MIGRATIONSCRIPTPATH string `mapstructure:"migrationscriptpath"`
}

func SetDefaults(v *viper.Viper) *viper.Viper {
	v.SetDefault("pg_user", "exoplanet_user")
	v.SetDefault("pg_pw", "password")
	v.SetDefault("pg_host", "localhost:5432")
	v.SetDefault("pg_db", "exoplanet")
	v.SetDefault("max_db_conns", "8")
	v.SetDefault("migrationscriptpath", "file://.db/migrations")
	v.SetDefault("environment", "dev")
	v.SetDefault("log_level", "debug")
	v.SetDefault("port", "8080")

	return v
}

func GetConfig() (Config, error) {
	v := viper.New()
	v.AutomaticEnv()
	v = SetDefaults(v)

	config := new(Config)
	if err := v.Unmarshal(config, func(cfg *mapstructure.DecoderConfig) {
		cfg.TagName = "mapstructure"
	}); err != nil {
		return *config, fmt.Errorf("config: error reading configuration into memory: %s", err)
	}

	return *config, nil
}
