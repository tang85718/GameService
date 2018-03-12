package gamemode

import (
	"github.com/tangxuyao/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func createNewActor(token string, name string) (*mongo.Charactor, error) {
	ms, err := mgo.Dial("")
	if err != nil {
		return nil, err
	}

	playerCol := ms.DB(mongo.DB_ROOT).C(mongo.C_PLAYER)
	player := mongo.Player{}
	err = playerCol.Find(bson.M{"token": token}).One(&player)
	if err != nil {
		return nil, err
	}

	actorCOL := ms.DB(mongo.DB_ROOT).C(mongo.C_ACTOR)
	count, err := actorCOL.Find(bson.M{"player_token": player.Token}).Count()

	if count > 0 {
		return nil, errors.New("不允许创建超过1个角色")
	}

	actor := mongo.Charactor{PlayerToken: token, Name: name, HP: 5, Energy: 0, EnergyType: 0}
	actorCOL.Insert(&actor)

	fmt.Printf("创建新角色%s, 属于玩家%s(%s)\n", name, player.DisplayID, player.Token)
	return &actor, nil
}