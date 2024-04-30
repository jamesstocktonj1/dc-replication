package main

import (
	"log"

	"github.com/jamesstocktonj1/dc-replication/app"
)

func main() {
	s := app.Server{}

	err := s.Init()
	if err != nil {
		log.Fatal(err)
	}

	s.Start()
	defer s.Stop()
}
