package main

import (
	"github.com/CherryRadiator/hakathon2025Spring/config"
	"github.com/CherryRadiator/hakathon2025Spring/server"
)

func main() {
	conf := config.GetConfig()
	s, err := server.NewGinServer(conf)
	if err != nil {
		panic(err)
	}

	s.RegisterHandlers(conf)	
	s.Start()
}
