package service

import "chess/model"

var Hub = &model.Hub{
	Rooms: make(map[model.RoomId]*model.Room),
	Ch:    make(chan *model.Message, 20),
}

var GameHub = &model.GameHub{
	Games: make(map[model.GameId]*model.Game),
	Ch:    make(chan *model.GameMessage, 20),
}
