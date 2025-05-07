package connector

import "time"

type DbConfig struct {
	URI            string        `mapstructure:"uri"`
	DB             string        `mapstructure:"db"`
	ConnectTimeout time.Duration `mapstructure:"connect_timeout"`
	PingTimeout    time.Duration `mapstructure:"ping_timeout"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}
