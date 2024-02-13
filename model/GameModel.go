package model

import (
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type WsPlayerInfo struct {
	Name string `json:"name"`
	Uid  string `json:"uid"`
	// player's info
	// can add more like rank
}

type GameRoom struct {
	ID        string                            `json:"id"`
	Players   map[*websocket.Conn]*WsPlayerInfo `json:"players"`
	GameState GameState                         `json:"game-state"`
	Settings  Settings                          `json:"settings"`
	StartTime time.Time                         `json:"start-time"`
	EndTime   time.Time                         `json:"end-time"`
	Lock      sync.Mutex
	Active    bool `json:"active"`
}

type GameState struct {
	Map         [][]Block               `json:"map"`
	Players     map[string]*PlayerState `json:"players"`
	CurrentTurn int                     `json:"current-turn"` // now Turn
	//	GameOver    bool                   `json:"game-over"`
	//	Winner      string                 `json:"winner"`
}

type Block struct {
	Type       int    `json:"type"`
	OccupiedBy string `json:"occupied-by"`
	UnitNumber int    `json:"unit-number"`
}

type PlayerState struct {
	Uid    string
	Name   string
	Color  string
	Units  int
	Occupy int
	Alive  bool
}

type Settings struct {
	Seed string `json:"seed"`
	// can add more like map size
}
