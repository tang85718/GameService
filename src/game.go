package gamemode

import (
	"github.com/micro/go-micro"
	"../../MongoData"
	"time"
)

type Game struct {
	M *mongo.MongoDB
}

func (g *Game) Run(service micro.Service) {
	defer func() {
		service.Server().Deregister()
	}()

	for {



		time.Sleep(time.Second * 5)
		//fmt.Println("[debug] Game Running..")
	}
}
