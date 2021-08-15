package config

import (
	"fmt"
	"line-wallet/db"

	"github.com/spf13/viper"
)

type Config struct {
	vn      *viper.Viper
	Server  Server     `mapstructure:"server"`
	MongoDB db.MongoDB `mapstructure:"mongodb"`
	LineBot LineBot    `mapstructure:"line-bot"`
}

type LineBot struct {
	ChannelAccessToken string `mapstructure:"channel-access-token"`
	ChannelSecret      string `mapstructure:"channel-secret"`
}

type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func (c *Config) InitAllConfiguration(env string) error {

	vn := viper.New()
	vn.SetConfigType("yaml")
	vn.SetConfigName(env)
	vn.AddConfigPath("../config/")
	vn.AddConfigPath("config/")
	c.vn = vn
	if err := vn.ReadInConfig(); err != nil {
		return fmt.Errorf("error on parsing configuration file")
	}

	if err := c.vn.Unmarshal(&c); err != nil {
		return fmt.Errorf("error Unmarshal %v", err)
	}

	if err := c.MongoDB.BindingClient(); err != nil {
		return fmt.Errorf("binding mongo error %v", err)

	}
	return nil
}

func (c *Config) GetConfig() *viper.Viper {
	return c.vn
}

func (c *Config) GetPort() string {
	return c.Server.Port
}

func (c *Config) GetHost() string {
	return c.Server.Host
}

func (c *Config) GetMongoUri() string {
	return c.MongoDB.MongoUri
}

func (c *Config) GetChannelSecret() string {
	return c.LineBot.ChannelSecret
}

func (c *Config) GetChannelAccessToken() string {
	return c.LineBot.ChannelAccessToken
}
