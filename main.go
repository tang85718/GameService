package main

import (
	"github.com/micro/go-micro"
	"proto/gm"
	"./src"
	//"fmt"
	//"strings"
	//consul "github.com/hashicorp/consul/api"
)

func main() {

	//config := consul.DefaultConfig()
	//config.Address = "localhost:8500"
	//c, err := consul.NewClient(config)
	//if err != nil {
	//	panic(err)
	//}
	//

	//ss, err := c.Agent().Services()
	//for k, v := range ss {
	//	fmt.Printf("%s", v)
	//
	//	//err = c.Agent().UpdateTTL(v.ID, "some bugs", consul.HealthPassing)
	//	//if err != nil {
	//	//	fmt.Printf(" %s\n", err)
	//	//} else {
	//	//	fmt.Println(" OK")
	//	//}
	//
	//	if strings.Contains(k, "game-mode") {
	//		c.Agent().ServiceDeregister(k)
	//	}
	//	//if err != nil {
	//	//	c.Agent().ServiceDeregister(k)
	//	//}
	//}

	service := micro.NewService(micro.Name("game-mode"))
	service.Init()

	defer func() {
		service.Server().Deregister()
	}()

	game := gamemode.Game{}
	game.Init()
	//go game.Push("one")
	//go game.Push("two")

	handler := new(gamemode.GameService)
	handler.Init(&game, service.Client())

	gm_api.RegisterGameServiceHandler(service.Server(), handler)
	go game.Run(service)

	service.Run()
}
