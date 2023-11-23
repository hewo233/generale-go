package model

type gameUser struct {
	Username string `json:"username"`
	Troopnum int    `json:"troopnum"`
	Rank     int    `json:"rank"`
	Color    string `json:"color"`
}
