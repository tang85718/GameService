package gamemode

import (
	"github.com/micro/go-micro"
	"./controllers"
	"fmt"
	"../../MongoData"
)

type Game struct {
	mgo *mongo.MongoDB
}

func (g *Game) Init() {
}

func (g *Game) execute(c controllers.Controller) {
}

func (g *Game) Run(service micro.Service) {
	defer func() {
		service.Server().Deregister()
	}()

	for {
		actor := mongo.Charactor{}
		if err := actor.FromDB(g.mgo); err != nil {

		}
		fmt.Println("[debug] Game Running..")
	}
}

func (g *Game) generateFirstTask() {
}
