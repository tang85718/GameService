package gamemode

import (
	"time"
	"container/list"
	"github.com/micro/go-micro"
	"fmt"
	"./controllers"
)

type Game struct {
	queue chan list.List
	token chan string
}

func (g *Game) Init() {
	g.token = make(chan string)
}

func (g *Game) Push(t string) {
	g.token <- t
}

func (g *Game) Run(service micro.Service) {
	defer func() {
		service.Server().Deregister()
	}()

	for {
		var input string
		input = <-g.token
		fmt.Printf("game running....%s\n", input)
		time.Sleep(time.Second)
		//panic("debug")
	}
}

func (g *Game) generateFirstTask() {
}
