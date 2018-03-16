package gamemode

import (
	"golang.org/x/net/context"
	"proto/gm"
	"../../MongoData"
	"proto/crm"
	"proto/asylum"
	"log"
)

type GameService struct {
	Crm     crm_api.CRMServiceClient
	Asylum  asylum_api.AsylumServiceClient
	factory *ActorFactory
}

func (s *GameService) Init(mgo *mongo.MongoDB) {
	s.factory = &ActorFactory{M: mgo}
}

func (s *GameService) StartGame(ctx context.Context, in *gm_api.StartGameReq, out *gm_api.StartGameRsp) error {
	// create actor
	actor, err := s.factory.createNew(in.Token, in.Name)
	if err != nil {
		log.Fatal(err)
		return err
	}

	actorID := actor.ID.Hex()

	_, err = s.Asylum.TakeActor(context.TODO(), &asylum_api.TakeActorReq{Token: actorID})
	if err != nil {
		return err
	}
	return nil
}

func (s *GameService) PingGame(c context.Context, in *gm_api.PingGameReq, out *gm_api.PingGameRsp) error {
	return nil
}
