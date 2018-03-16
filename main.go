package main

import (
	"github.com/micro/go-micro"
	"proto/gm"
	"./src"
	"../MongoData"
	"fmt"
	"proto/asylum"
	"proto/crm"
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

	cli := service.Client()
	crm := crm_api.NewCRMServiceClient("crmService", cli)
	asylum := asylum_api.NewAsylumServiceClient("AsylumService", cli)

	game := gamemode.Game{M: mg, Asylum: asylum}

	handler := &gamemode.GameService{Crm: crm, Asylum: asylum}
	handler.Init(mg)

	gm_api.RegisterGameServiceHandler(service.Server(), handler)

	fmt.Println("GameService starting...")

	go game.Run(service)
	service.Run()
}
