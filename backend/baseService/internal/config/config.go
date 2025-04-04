package config

import "sync"

type Config struct {
	Base
	Data
	Server
	Snowflake
}

var c *Config
var once sync.Once

func Init() {
	once.Do(func() {
		c = new(Config)
	})
}

func GetConfig() *Config {
	if c == nil {
		once.Do(func() {
			c = new(Config)
		})
	}
	return c
}
