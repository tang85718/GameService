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