package http

import (
	"fmt"
	"time"
)

type Config struct {
	Host         string
	Port         uint16
	Silent       bool
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func (c Config) DNS() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c Config) Path() string {
	return fmt.Sprintf("http://%s:%d", c.Host, c.Port)
}
