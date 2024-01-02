package main

import "fmt"

type Server struct {
	config Config
}

func NewServer(config Config) *Server {
	return &Server{
		config: config,
	}
}

type Config struct {
	iD            string
	name          string
	listenAddress string
}

func (c Config) WithID(id string) Config {
	c.iD = id
	return c
}

func (c Config) WithName(name string) Config {
	c.name = name
	return c
}

func (c Config) WithListenAddr(listenAddr string) Config {
	c.listenAddress = listenAddr
	return c
}

func NewConfig() Config {
	return Config{
		iD:            "default_id",
		name:          "default_name",
		listenAddress: "3000",
	}
}

func main() {
	config := NewConfig().
		WithID("a4554dasd54dsa").
		WithName("super_server").
		WithListenAddr("8000")
	server := NewServer(config)

	fmt.Println(server)
}
