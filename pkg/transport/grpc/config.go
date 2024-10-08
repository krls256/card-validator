package grpc

import "fmt"

type Config struct {
	Host string
	Port uint16
}

func (c Config) DNS() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
