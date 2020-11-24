package main

import (
	"flag"
	"fmt"
	"os"
	"seminario-GoLang/internal/config"
	"seminario-GoLang/internal/service/instruments"
)

func main() {
	cfg := readConfig()

	service, _ := instruments.New(cfg)
	for _, m := range service.FindAll() {
		fmt.Println(m)
	}
}

func readConfig() *config.Config {
	configFile := flag.String("config", "./config.yaml", "this is how you should config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return cfg
}
