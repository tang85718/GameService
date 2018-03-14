package controllers

import (
	"../../../MongoData"
	"fmt"
)

type ActorController struct {
	actor *mongo.Charactor
	mgo   *mongo.MongoDB
}

func (ac *ActorController) Run() error {
	fmt.Printf("run actor controller %s\n", ac.actor.PlayerToken)
	return nil
}

func (ac *ActorController) CreateTODO() error {
	token := ac.actor.ID.Hex()
	db := ac.mgo.Conn.DB(token).C(mongo.C_TODO)
	count, err := db.Find(nil).Count()
	if err != nil {
		return err
	}

	if count == 0 {

	}

	todo := mongo.Todo{
		ActorToken: ac.actor.ID.Hex(),
	}
}
