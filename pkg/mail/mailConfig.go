package mail_config

import (
	"fmt"

	"github.com/spf13/viper"
)

type MAILCfg struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func ReadConfigMAIL() (*MAILCfg, error) {
	viper.SetConfigFile("internal/config/mail.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %s", err)
	}

	var cfg MAILCfg

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %s", err)
	}

	return &cfg, nil
}
