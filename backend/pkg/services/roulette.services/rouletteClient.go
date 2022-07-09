package rouletteservices

import "github.com/gorilla/websocket"

type RouletteClient struct {
	ClientId string `json:"id"`
	WsConn   *websocket.Conn
}

func (rClient *RouletteClient) ReadRouletteClient(rGame *RouletteGame) {
	defer func() {
		rGame.UnregisterPlayerConn <- rClient
		rClient.WsConn.Close()
	}()

}
