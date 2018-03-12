package gamemode

import (
	"github.com/tangxuyao/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"errors"
	"fmt"
)

type ActorFactory struct {
	ms *mgo.Session
}

func (f *ActorFactory) Init() error {
	ms, err := mgo.Dial("")
	if err != nil {
		return err
	}
	f.ms = ms
	return nil
}

func (f *ActorFactory) CheckHealth() error {
	if f.ms.Ping() != nil {
		return nil
	}

	f.ms.Close()
	f.Init()

	return nil
}

func (f *ActorFactory) createNewActor(token string, name string) (*mongo.Charactor, error) {
	playerCol := f.ms.DB(mongo.DB_ROOT).C(mongo.C_PLAYER)
	player := mongo.Player{}
	err := playerCol.Find(bson.M{"token": token}).One(&player)
	if err != nil {
		return nil, err
	}

	actorCOL := f.ms.DB(mongo.DB_ROOT).C(mongo.C_ACTOR)
	count, err := actorCOL.Find(bson.M{"player_token": player.Token}).Count()

	if count > 0 {
		return nil, errors.New("不允许创建超过1个角色")
	}

	actor := mongo.Charactor{PlayerToken: token, Name: name, HP: 5, Energy: 0, EnergyType: 0}
	actorCOL.Insert(&actor)

	fmt.Printf("创建新角色%s, 属于玩家%s(%s)\n", name, player.DisplayID, player.Token)
	return &actor, nil
}
