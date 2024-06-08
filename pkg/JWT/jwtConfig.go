package jwt

import (
	"fmt"

	"github.com/spf13/viper"
)

type JWTCfg struct {
	SecretKey string `mapstructure:"secretKey"`
	Mdtd      string `mapstructure:"method"`
}

func ReadConfigJWT() (*JWTCfg, error) {
	viper.SetConfigFile("internal/config/jwt.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %s", err)
	}

	var cfg JWTCfg

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %s", err)
	}

	return &cfg, nil
}
