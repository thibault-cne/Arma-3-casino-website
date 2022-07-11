package rouletteservices

import (
	"fmt"
	"time"

	"casino.website/pkg/models"
)

type RouletteGame struct {
	Id                   string
	RegisterBet          chan *models.RouletteBets
	RegisterPlayerConn   chan *RouletteClient
	UnregisterPlayerConn chan *RouletteClient
	Bets                 map[*models.RouletteBets]bool
	ConnectedPlayers     map[*RouletteClient]bool
	BroadcastBets        chan *models.RouletteBets
}

func NewRouletteGame() *RouletteGame {
	return &RouletteGame{
		Id:                   generateRandomId(),
		RegisterBet:          make(chan *models.RouletteBets),
		RegisterPlayerConn:   make(chan *RouletteClient),
		UnregisterPlayerConn: make(chan *RouletteClient),
		Bets:                 make(map[*models.RouletteBets]bool),
		ConnectedPlayers:     make(map[*RouletteClient]bool),
		BroadcastBets:        make(chan *models.RouletteBets),
	}
}

func (rGame *RouletteGame) cleanRouletteGame() {
	rGame.Bets = make(map[*models.RouletteBets]bool)
}

func (rGame *RouletteGame) Start() {
	for {
		select {
		case bet := <-rGame.RegisterBet:
			fmt.Printf("Registering a new bet %+v\n", bet)
			rGame.Bets[bet] = true
			break
		case client := <-rGame.RegisterPlayerConn:
			fmt.Printf("Registering a new player %+v\n", client)
			rGame.ConnectedPlayers[client] = true
			break
		case client := <-rGame.UnregisterPlayerConn:
			delete(rGame.ConnectedPlayers, client)
			break
		case bet := <-rGame.BroadcastBets:
			for connectedClient := range rGame.ConnectedPlayers {
				if err := connectedClient.WsConn.WriteJSON(bet); err != nil {
					fmt.Printf("An error occured while broadcasting : %s", err.Error())
					return
				}
			}
			break
		}
	}
}

func (rGame *RouletteGame) End() {
	defer func() {
		fmt.Printf("Ended a roulette game.\n")
		rGame.cleanRouletteGame()
		fmt.Printf("Game cleaned.\n")
		go rGame.End()
	}()

	time.Sleep(5 * time.Second)

	color := calculateRouletteColor()

	fmt.Printf("The good color is : %s\n", color)

	for bet := range rGame.Bets {
		if bet.BetPosition == color {
			fmt.Printf("Client %s won %d\n", bet.ClientId, bet.BetAmount)
		}
	}
}
