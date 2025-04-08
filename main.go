package main

import (
	"github.com/dmxmss/news-service/server"
	"github.com/dmxmss/news-service/config"
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
