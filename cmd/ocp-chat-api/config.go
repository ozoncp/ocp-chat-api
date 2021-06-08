package main

const defaultTransportProtocol = "tcp"

type Config struct {
	HTTPAddr string `envconfig:"HTTP_ADDR"`
	GRPCAddr string `envconfig:"GRPC_ADDR"`
}

func NewDefaultConfig() *Config {
	return &Config{
		HTTPAddr: ":8080",
		GRPCAddr: ":5300",
	}
}
