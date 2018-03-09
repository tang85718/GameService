package gamemode

import (
	"time"
	"fmt"
	"gopkg.in/mgo.v2"
	"github.com/tangxuyao/mongo"
)

type Game struct {
	actors []mongo.Charactor
}

func (g *Game) Run() {

	m, err := mgo.Dial("")
	if err != nil {
		panic(err)
	}

	for {
		/**
		拿到所有角色数据，将数据放入RabbitMQ队列之中

		*/
		if len(g.actors) == 0 {
			col := m.DB("crm").C("actors")
			col.Find(nil).All(&g.actors)
		}

		// 如何判断是新出生的角色

		for i := 0; i < 100; i++ {
			actor := g.actors[i]
			todoCol := m.DB(actor.PlayerToken).C("todo")

			todoCount, _ := todoCol.Count()
			if todoCount == 0 {
				g.generateFirstTask()
			}
		}

		fmt.Println("game running....")
		time.Sleep(time.Second)
	}
}

func (g *Game) generateFirstTask() {
}
