package gamemode

import (
	"golang.org/x/net/context"
	"proto/gm"
	"github.com/micro/go-micro/client"
	"proto/crm"
	"./controllers"
)

type GameService struct {
	crm     crm_api.CRMServiceClient
	game    *Game
	factory *ActorFactory
}

func (s *GameService) Init(g *Game, c client.Client) {
	s.crm = crm_api.NewCRMServiceClient("crmService", c)
	s.factory = &ActorFactory{g.mgo}
	s.game = g
}

func (s *GameService) StartGame(c context.Context, in *gm_api.StartGameReq, out *gm_api.SimpleRsp) error {
	// create actor
	n, err := s.factory.createNew(in.Token, in.Name)
	if err != nil {
		return err
	}

	factory := controllers.ControllerFactory{}
	ac := factory.CreateController(s.game.mgo, n)
	ac.CreateTODO()
	return nil
}
