package service

import "chess/model"

var Hub = &model.Hub{
	Rooms: make(map[model.RoomId]*model.Room),
	Games: make(map[model.GameId]*model.Game),
	Ch:    make(chan *model.Message, 20),
}
