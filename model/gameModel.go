package model

type gameConfig struct {
	Blocks           []string               `json:"blocks"`
	EventTypes       []string               `json:"eventTypes"`
	EventBinaryChunk EventBinaryChunkDetail `json:"event-binary-chunk"`
	EventJSON        EventJSONDetail        `json:"event-json"`
}

type EventBinaryChunkDetail struct {
	GameUpdateMap         GameUpdateMapDetail         `json:"game-update-map"`
	GameUpdateActionQueue GameUpdateActionQueueDetail `json:"game-update-action-queue"`
	GameFullMap           GameFullMapDetail           `json:"game-full-map"`
}

type GameUpdateMapDetail struct {
	Length      int                 `json:"length"`
	BlocksBegin int                 `json:"blocksBegin"`
	Blocks      GameUpdateMapBlocks `json:"blocks"`
}

type GameUpdateMapBlocks struct {
	Row       int `json:"row"`
	Col       int `json:"col"`
	Id        int `json:"id"`
	PlayedId  int `json:"played-id"`
	Number    int `json:"number"`
	BlockSize int `json:"block-size"`
}

type GameUpdateActionQueueDetail struct {
	Length     int                         `json:"length"`
	QueueBegin int                         `json:"queueBegin"`
	Action     GameUpdateActionQueueAction `json:"action"`
}

type GameUpdateActionQueueAction struct {
	Row       int `json:"row"`
	Col       int `json:"col"`
	Act       int `json:"act"`
	BlockSize int `json:"block-size"`
}

type GameFullMapDetail struct {
	Rows        int               `json:"rows"`
	Cols        int               `json:"cols"`
	BlocksBegin int               `json:"blocksBegin"`
	Blocks      GameFullMapBlocks `json:"blocks"`
}

type GameFullMapBlocks struct {
	Id        int `json:"id"`
	PlayedId  int `json:"played-id"`
	Number    int `json:"number"`
	BlockSize int `json:"block-size"`
}

type EventJSONDetail struct {
	WsOnOpen              []interface{}         `json:"ws-onopen"`
	WsOnClose             []interface{}         `json:"ws-onclose"`
	RoomUpdateInfos       RoomUpdateInfos       `json:"room-update-infos"`
	ChatboxMessageReceive ChatboxMessageReceive `json:"chatbox-message-receive"`
	GameStart             []interface{}         `json:"game-start"`
	GameRequestMove       GameRequestMove       `json:"game-request-move"`
	GameRequestClearMove  []interface{}         `json:"game-request-clear-move"`
	GameRequestFullMap    []interface{}         `json:"game-request-full-map"`
	GameUpdateRank        GameUpdateRank        `json:"game-update-rank"`
	GameOver              []interface{}         `json:"game-over"`
}

type RoomUpdateInfos struct {
	Players map[string]PlayerInfo `json:"players"`
}

type PlayerInfo struct {
	Color string `json:"color"`
	UUID  string `json:"uuid"`
}

type ChatboxMessageReceive struct {
	ID      int    `json:"id"`
	Time    string `json:"time"`
	Message string `json:"message"`
}

type GameRequestMove struct {
	Row int    `json:"row"`
	Col int    `json:"col"`
	Act string `json:"act"`
}

type GameUpdateRank struct {
	Rank []PlayerRank `json:"rank"`
}

type PlayerRank struct {
	ID    int `json:"id"`
	Army  int `json:"army"`
	Lands int `json:"lands"`
}
