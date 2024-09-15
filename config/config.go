package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	EnvironmentVariables EnvironmentVariables `yaml:"EnvironmentVariables"`
	ConnectionStrings    ConnectionStrings    `yaml:"ConnectionStrings"`
	HTTPServer           HTTPServer           `yaml:"HTTPServer"`
}

type EnvironmentVariables struct {
	Environment string `yaml:"Environment"`
}

type ConnectionStrings struct {
	ServiceDb string `yaml:"ServiceDb"`
}

type HTTPServer struct {
	Url string `yaml:"Url"`
}

func ReadConfig(cfgName, cfgType, cfgPath string) (*Config, error) {
	var cfg Config

	viper.SetConfigName(cfgName)
	viper.SetConfigType(cfgType)
	viper.AddConfigPath(cfgPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
