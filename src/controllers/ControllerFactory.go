package controllers

import "../../../MongoData"

type ControllerFactory struct {
}

func (f *ControllerFactory) CreateController(mgo *mongo.MongoDB, c *mongo.Charactor) ActorController {
	a := ActorController{actor: c, mgo: mgo}
	return a
}
