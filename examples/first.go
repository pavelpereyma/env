package main

import (
	"fmt"

	"github.com/caarlos0/env"
)

type config struct {
	Home         string `env:"HOME"`
	Port         int    `env:"PORT"`
	IsProduction bool   `env:"PRODUCTION"`
}

func main() {
	cfg := config{}
	env.Parse(&cfg)
	fmt.Println(cfg)
}