package gamemode

import (
	"github.com/micro/go-micro"
	"../../MongoData"
	"time"
	"proto/asylum"
	"golang.org/x/net/context"
	"log"
)

type Game struct {
	M      *mongo.MongoDB
	Asylum asylum_api.AsylumServiceClient
}

func (g *Game) Run(service micro.Service) {
	defer func() {
		service.Server().Deregister()
	}()

	db := g.M.Conn.DB(mongo.DB_GLOBAL).C(mongo.C_ACTOR)
	itr := db.Find(nil).Iter()

	for {
		actor := mongo.Charactor{}
		if itr.Next(&actor) {
			switch actor.Place {
			case mongo.PLACE_GOD_SPACE:
				_, err := g.Asylum.StartStory(context.TODO(), &asylum_api.StartStoryReq{Token: actor.ID.Hex()})
				if err != nil {
					log.Println("Invoke StartStory error: " + err.Error())
				}
			case mongo.PLACE_ASYLUM:
				_, err := g.Asylum.RunStory(context.TODO(), &asylum_api.RunStoryReq{Token: actor.ID.Hex()})
				if err != nil {
					log.Fatalln("Invoke RunStory error: " + err.Error())
				}
			case mongo.PLACE_WILDERNESS:
			default:
			}
		}

		time.Sleep(time.Second * 5)
	}
}
