package controllers

import "github.com/tangxuyao/mongo"

type ActorController struct {
	actor *mongo.Charactor
}

func (a *ActorController) Run() error {
	return nil
}
