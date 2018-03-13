package controllers

import "github.com/tangxuyao/mongo"

type ControllerFactory struct {
}

func (f *ControllerFactory) CreateController(c *mongo.Charactor) ActorController {
	a := ActorController{actor:c}
	return &a
}
