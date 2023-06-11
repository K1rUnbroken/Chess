package service

import "chess/model"

var HubListener = &model.Hub{
	Rooms: make(map[int]*model.Room),
	Srv:   make(chan *model.Message, 20),
}
