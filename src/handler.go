package gamemode

import (
	"golang.org/x/net/context"
	"proto/gm"
	"github.com/micro/go-micro/client"
	"proto/crm"
)

type GameService struct {
	crm crm_api.CRMServiceClient
}

func (s *GameService) Init(c client.Client) {
	s.crm = crm_api.NewCRMServiceClient("crmService", c)
}

func (s *GameService) StartGame(c context.Context, in *gm_api.StartGameReq, out *gm_api.SimpleRsp) error {

	// create actor
	_, err := s.crm.MakeActor(context.TODO(), &crm_api.MakeActorReq{Token: in.Token, Name: in.Name})
	if err != nil {
		return err
	}

	return nil
}
