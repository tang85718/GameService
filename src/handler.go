package gamemode

import (
	"golang.org/x/net/context"
	"proto/gm"
	"github.com/micro/go-micro/client"
	"proto/crm"
	"fmt"
	"./activities"
)

type GameService struct {
	crm     crm_api.CRMServiceClient
	game    *Game
	factory *ActorFactory
}

func (s *GameService) Init(g *Game, c client.Client) {
	s.crm = crm_api.NewCRMServiceClient("crmService", c)
	s.game = g
	s.factory = new(ActorFactory)
	s.factory.Init()
}

func (s *GameService) StartGame(c context.Context, in *gm_api.StartGameReq, out *gm_api.SimpleRsp) error {
	go s.game.Push(in.Token)

	// create actor
	n, err := s.factory.createNewActor(in.Token, in.Name)
	if err != nil {
		return err
	}

	fmt.Println(n.PlayerToken)
	//time.Sleep(time.Second * 3)
	return nil
}
