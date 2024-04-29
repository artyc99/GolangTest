package main

import (
	"EchoAPI/infrastructure/configs"
	"EchoAPI/infrastructure/controllers/http_server"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the HELL!")

	cfg, err := configs.GetHttpConfig()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cfg)

	err = http_server.RunServer(cfg)
	if err != nil {
		fmt.Println(err)
	}
}
