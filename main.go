package main

import (
	"github.com/micro/go-micro"
	"proto/gm"
	"./src"
)

func main() {
	service := micro.NewService(micro.Name("game-mode"))
	service.Init()

	defer func() {
		service.Server().Deregister()
	}()

	game := gamemode.Game{}
	game.Init()

	handler := new(gamemode.GameService)
	handler.Init(&game, service.Client())

	gm_api.RegisterGameServiceHandler(service.Server(), handler)

	go game.Run(service)
	service.Run()
}
