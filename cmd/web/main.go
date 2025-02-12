package main

import (
	"flag"
	"github.com/robertgouveia/portfolio/internal/server"
)

func main() {
	var env string
	var port string
	flag.StringVar(&env, "env", "development", "environment to use")
	flag.StringVar(&port, "port", ":443", "port to use")
	flag.Parse()

	s := server.NewServer(port, env)

	go s.ListenForShutdown()

	if err := s.Run(); err != nil {
		panic(err)
	}
}
