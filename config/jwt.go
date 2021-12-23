package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

//JTWConfig
type JWTConfig struct {
	Secret         string `json:"secret"`
	ExpirationTime int    `json:"expiration-time"`
}

var (
	jwtConfig = JWTConfig{
		Secret:         "MySecret",
		ExpirationTime: 1,
	}
)

//GetJwtConf return the parameters for jwt conf
func GetJwtConf() (JWTConfig, error) {
	if err := mapstructure.Decode(
		viper.GetStringMap("jwt"),
		&jwtConfig,
	); err != nil {
		return JWTConfig{}, err
	}
	return jwtConfig, nil
}
