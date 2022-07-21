package main

import (
	"flag"
	"fmt"

	"publish/client"
	"publish/config"
	"publish/server"
)

func main() {
	var (
		path string
	)

	flag.StringVar(&path, "config", "./default.config.toml", `config file path (*.config.toml")`)
	flag.Parse()
	// fmt.Println(path)

	if err := config.LoadConfigFile(path); err != nil {
		fmt.Printf("load config file error: %v\n", err)
		config.LoadDefaultConfig()
	}

	if config.C.Server.Debug {
		fmt.Printf("%+v\n", config.C)
	}

	client.InitClient()

	router := server.Server()
	router.Run(fmt.Sprintf(":%d", config.C.Server.Port))
}
