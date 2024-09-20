package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	EnvironmentVariables EnvironmentVariables `yaml:"EnvironmentVariables" json:"EnvironmentVariables"`
	ConnectionStrings    ConnectionStrings    `yaml:"ConnectionStrings" json:"ConnectionStrings"`
	HTTPServer           HTTPServer           `yaml:"HTTPServer" json:"HTTPServer"`
	Logs                 Logs                 `yaml:"Logs" json:"Logs"`
}

type EnvironmentVariables struct {
	Environment string `yaml:"Environment" json:"Environment"`
}

type ConnectionStrings struct {
	ServiceDb string `yaml:"ServiceDb" json:"ServiceDb"`
}

type HTTPServer struct {
	Addr string `yaml:"Addr" json:"Addr"`
	Port int    `yaml:"Port" json:"Port"`
}

type Logs struct {
	Path       string `yaml:"Path" json:"Path"`
	Level      string `yaml:"Level" json:"Level"`
	MaxAge     int    `yaml:"MaxAge" json:"MaxAge"` // measured in Days
	MaxBackups int    `yaml:"MaxBackups" json:"MaxBackups"`
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
