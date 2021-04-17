package main

import (
	"log"

	"github.com/Baccanimonium/goRestApi/internal/app/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
