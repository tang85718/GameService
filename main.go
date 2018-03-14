package main

import (
	"github.com/micro/go-micro"
	"proto/gm"
	"./src"
	"../MongoData"
)

func main() {
	mg := &mongo.MongoDB{}
	if err := mg.Dial(""); err != nil {
		panic("连接[mongo]数据失败，请检查相关参数")
	}

	service := micro.NewService(micro.Name("game-mode"))
	service.Init()

	defer func() {
		service.Server().Deregister()
	}()

	game := gamemode.Game{mg}
	game.Init()

	handler := new(gamemode.GameService)
	handler.Init(&game, service.Client())

	gm_api.RegisterGameServiceHandler(service.Server(), handler)

	go game.Run(service)
	service.Run()
}
