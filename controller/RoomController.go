package controller

import (
	"generale-go/model"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type GameRoom struct {
	model.GameRoom
}

func CreatNewRoom(roomID string, newSettings model.Settings) *model.GameRoom {
	room := &model.GameRoom{
		ID:      roomID,
		Players: make(map[*websocket.Conn]*model.WsPlayerInfo),
		GameState: model.GameState{
			Players: make(map[string]*model.PlayerState),
		},
		Settings:  newSettings,
		StartTime: time.Now(),
		EndTime:   time.Now(),
		Lock:      sync.Mutex{},
		Active:    false,
	}
	return room
}

func (r *GameRoom) NewPlayerConnect(conn *websocket.Conn, playerUid string, playerName string) {

	r.Lock.Lock()
	defer r.Lock.Unlock()

	newWsPlayer := &model.WsPlayerInfo{
		Uid:  playerUid,
		Name: playerName,
	}

	r.Players[conn] = newWsPlayer

	newStatePlayer := &model.PlayerState{
		Uid:    playerUid,
		Name:   playerName,
		Color:  GenerateColor(playerUid),
		Units:  1,
		Occupy: 1,
		Alive:  true,
	}

	r.GameState.Players[playerUid] = newStatePlayer

	r.Active = true

}

func (r *GameRoom) PlayerDisconnect(conn *websocket.Conn) {
	r.Lock.Lock()
	defer r.Lock.Unlock()

	if wsPlayer, exits := r.Players[conn]; exits {
		delete(r.GameState.Players, wsPlayer.Uid)
		delete(r.Players, conn)
	}
}
