package rouletteservices

import (
	"encoding/json"
	"fmt"
	"log"

	"casino.website/pkg/models"
	"github.com/gorilla/websocket"
)

type RouletteClient struct {
	ClientId string `json:"id"`
	WsConn   *websocket.Conn
}

func (rClient *RouletteClient) ReadRouletteClient(rGame *RouletteGame) {
	defer func() {
		fmt.Printf("Defering player conn\n")
		rGame.UnregisterPlayerConn <- rClient.WsConn
		rClient.WsConn.Close()
	}()

	for {
		_, p, err := rClient.WsConn.ReadMessage()
		fmt.Printf("Receiving new message ...\n")
		if err != nil {
			log.Println(err)
			return
		}

		var frontBet models.FrontBet

		if err = json.Unmarshal(p, &frontBet); err != nil {
			fmt.Printf("An error occured while unmarshalling data : %s\n", err.Error())
			fmt.Printf("Here is what where sent : %q\n", p)
			return
		}

		rBet := &models.RouletteBets{ClientId: rClient.ClientId, BetPosition: frontBet.Color, BetAmount: frontBet.Amount}

		fmt.Printf("%p", rBet)

		rGame.RegisterBet <- rBet
		rGame.BroadcastBets <- rBet
	}
}
