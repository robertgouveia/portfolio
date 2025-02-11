package main

import (
	"flag"
	"github.com/robertgouveia/portfolio/internal/server"
)

func main() {
	var env string
	flag.StringVar(&env, "env", "development", "environment to use")
	flag.Parse()

	s := server.NewServer(":80", env)

	go s.ListenForShutdown()

	if err := s.Run(); err != nil {
		panic(err)
	}
}
