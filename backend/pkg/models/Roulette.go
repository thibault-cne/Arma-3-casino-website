package models

import "github.com/gorilla/websocket"

type RouletteBets struct {
	ClientId    string `json:"id"`
	BetPosition string `json:"position"`
	BetAmount   int    `json:"amount"`
}

type RouletteClient struct {
	ClientId string `json:"id"`
	WsConn   *websocket.Conn
}
