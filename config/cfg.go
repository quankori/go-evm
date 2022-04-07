package config

import "github.com/spf13/viper"

type Config struct {
	PrivateKey    string `mapstructure:"PRIVATE_KEY"`
	NetworkURI    string `mapstructure:"NETWORK_URI"`
	WebsocketkURI string `mapstructure:"WEBSOCKET_URI"`
	TokenURI      string `mapstructure:"TOKEN_URI"`
}

func LoadConfig() (config Config, err error) {
	viper.AutomaticEnv()
	viper.AddConfigPath("./")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
