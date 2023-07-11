package main

import (
	"log"

	"github.com/suhas-totality/gomongodemo/pkg/env"
)

var preInit *Init

func init() {
	envPlain, err := env.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}
	preInit = newInit(envPlain)
}
