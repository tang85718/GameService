package main

import (
	"github.com/micro/go-micro"
	"proto/gm"
	"./src"
	"../MongoData"
	"fmt"
)

func main() {
	mg := &mongo.MongoDB{}
	if err := mg.Dial(""); err != nil {
		panic("连接[mongo]数据失败，请检查相关参数")
	}

	service := micro.NewService(micro.Name("GameService"))
	service.Init()

	defer func() {
		service.Server().Deregister()
	}()

	game := gamemode.Game{M: mg}

	handler := new(gamemode.GameService)
	handler.Init(&game, service.Client())

	gm_api.RegisterGameServiceHandler(service.Server(), handler)

	fmt.Println("GameService starting...")

	go game.Run(service)
	service.Run()
}
