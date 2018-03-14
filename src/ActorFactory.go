package gamemode

import (
	"../../MongoData"
	"fmt"
)

type ActorFactory struct {
	mgo *mongo.MongoDB
}

func (f *ActorFactory) createNew(token string, name string) (*mongo.Charactor, error) {
	player := mongo.Player{}
	err := player.FromDB(f.mgo, token)
	if err != nil {
		return nil, err
	}

	//todo: 添加环境开关，debug模式下不运行下面代码
	//count, err := actorCOL.Find(bson.M{"player_token": player.Token}).Count()
	//if count > 0 {
	//	return nil, errors.New("不允许创建超过1个角色")
	//}

	actor := mongo.Charactor{
		PlayerToken: token,
		Name:        name,
		HP:          5,
		Energy:      0,
		EnergyType:  0,
	}

	actor.ToDB(f.mgo)

	fmt.Printf("创建新角色%s, 属于玩家%s(%s)\n", name, player.DisplayID, player.Token)
	return &actor, nil
}
