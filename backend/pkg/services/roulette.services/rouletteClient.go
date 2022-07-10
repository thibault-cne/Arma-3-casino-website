package rouletteservices

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type RouletteClient struct {
	ClientId string `json:"id"`
	WsConn   *websocket.Conn
}

func (rClient *RouletteClient) ReadRouletteClient(rGame *RouletteGame) {
	defer func() {
		rGame.UnregisterPlayerConn <- rClient
		rClient.WsConn.Close()
	}()

	for {
		messageType, p, err := rClient.WsConn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("Message Received: %d || %+v\n", messageType, p)
	}
}
