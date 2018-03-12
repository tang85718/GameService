package gamemode

import (
	"time"
	"fmt"
	"gopkg.in/mgo.v2"
	"github.com/tangxuyao/mongo"
	"container/list"
)

type Game struct {
	actors list.List
}

func (g *Game) Run() {

	ms, err := mgo.Dial("")
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("game running....")
		time.Sleep(time.Second)
	}
}

func (g *Game) generateFirstTask() {
}
