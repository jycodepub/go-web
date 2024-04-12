package config

import (
	"fmt"
	"os"
	"strconv"
)

const (
	DefaultPort         = 8080
	DefaultWriteTimeout = 15
	DefaultReadTimeout  = 15
)

type Config struct {
	Port         int
	WriteTimeout int
	ReadTimeout  int
}

func (c *Config) Address() string {
	return fmt.Sprintf(":%d", c.Port)
}

func GetIntValue(key string, defaultValue int) int {
	value, found := os.LookupEnv(key)
	if !found {
		return defaultValue
	} else {
		v, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue
		}
		return v
	}
}

func GetConfig() *Config {
	port := GetIntValue("PORT", DefaultPort)
	wt := GetIntValue("WRITE_TIMEOUT", DefaultWriteTimeout)
	rt := GetIntValue("READ_TIMEOUT", DefaultReadTimeout)
	return &Config{
		Port:         port,
		WriteTimeout: wt,
		ReadTimeout:  rt,
	}
}
