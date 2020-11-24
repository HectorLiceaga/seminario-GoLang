package main

import (
	"flag"
	"fmt"
	"os"
	"seminario-GoLang/internal/config"
)

func main() {
	configFile := flag.String("config", "./config.yaml", "this is how you should config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
