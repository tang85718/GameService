package main

import (
	"github.com/micro/go-micro"
	"proto/gm"
	"./src"
)

func main() {
	service := micro.NewService(micro.Name("game-mode"))
	service.Init()

	handler := new(gamemode.GameService)
	handler.Init(service.Client())

	gm_api.RegisterGameServiceHandler(service.Server(), handler)

	go service.Run()

	game := gamemode.Game{}
	game.Run()
}
