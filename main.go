package main

import (
	"github.com/micro/go-micro"
	"proto/gm"
	"./src"
)

func main() {
	service := micro.NewService(micro.Name("game-mode"))
	service.Init()

	gm := new(gamemode.GameService)
	gm.Init(service.Client())

	gm_api.RegisterGameServiceHandler(service.Server(), gm)

	go service.Run()

	game := gamemode.Game{}
	game.Run()
}
