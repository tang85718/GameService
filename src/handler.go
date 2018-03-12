package gamemode

import (
	"golang.org/x/net/context"
	"proto/gm"
	"github.com/micro/go-micro/client"
	"proto/crm"
	"fmt"
	"time"
)

type GameService struct {
	crm  crm_api.CRMServiceClient
	game *Game
}

func (s *GameService) Init(g *Game, c client.Client) {
	s.crm = crm_api.NewCRMServiceClient("crmService", c)
	s.game = g
}

func (s *GameService) StartGame(c context.Context, in *gm_api.StartGameReq, out *gm_api.SimpleRsp) error {
	fmt.Println("Starting Game...")
	// create actor
	createNewActor(in.Token, in.Name)
	time.Sleep(time.Second * 3)
	return nil
}
