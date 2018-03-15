package gamemode

import (
	"golang.org/x/net/context"
	"proto/gm"
	"github.com/micro/go-micro/client"
	"proto/crm"
	"proto/asylum"
	"log"
)

type GameService struct {
	crm     crm_api.CRMServiceClient
	asylum  asylum_api.AsylumServiceClient
	game    *Game
	factory *ActorFactory
}

func (s *GameService) Init(g *Game, c client.Client) {
	s.crm = crm_api.NewCRMServiceClient("crmService", c)
	s.asylum = asylum_api.NewAsylumServiceClient("AsylumService", c)
	s.factory = &ActorFactory{g.M}
	s.game = g
}

func (s *GameService) StartGame(ctx context.Context, in *gm_api.StartGameReq, out *gm_api.StartGameRsp) error {
	log.Printf("Player Token:%s\n", in.Token)
	// create actor
	actor, err := s.factory.createNew(in.Token, in.Name)
	if err != nil {
		log.Fatal(err)
		return err
	}

	actorID := actor.ID.Hex()

	r := asylum_api.TakeActorReq{Token: actorID}
	_, err = s.asylum.TakeActor(context.Background(), &r)
	if err != nil {
		return err
	}

	log.Println("Start Game Successful!!!")
	return nil
}

func (s *GameService) PingGame(c context.Context, in *gm_api.PingGameReq, out *gm_api.PingGameRsp) error {
	return nil
}
