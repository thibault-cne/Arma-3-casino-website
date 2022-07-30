package rouletteservices

import (
	"fmt"
	"time"

	"casino.website/pkg/models"
	"github.com/gorilla/websocket"
)

type RouletteGame struct {
	Id                   string
	RegisterBet          chan *models.RouletteBets
	RegisterPlayerConn   chan *websocket.Conn
	UnregisterPlayerConn chan *websocket.Conn
	Bets                 map[*models.RouletteBets]bool
	ConnectedPlayers     map[*websocket.Conn]bool
	BroadcastBets        chan *models.RouletteBets
}

func NewRouletteGame() *RouletteGame {
	return &RouletteGame{
		Id:                   generateRandomId(),
		RegisterBet:          make(chan *models.RouletteBets),
		RegisterPlayerConn:   make(chan *websocket.Conn),
		UnregisterPlayerConn: make(chan *websocket.Conn),
		Bets:                 make(map[*models.RouletteBets]bool),
		ConnectedPlayers:     make(map[*websocket.Conn]bool),
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
		case conn := <-rGame.RegisterPlayerConn:
			fmt.Printf("Registering a new player.\n")
			rGame.ConnectedPlayers[conn] = true
			break
		case conn := <-rGame.UnregisterPlayerConn:
			delete(rGame.ConnectedPlayers, conn)
			break
		case bet := <-rGame.BroadcastBets:
			for conn := range rGame.ConnectedPlayers {
				if err := conn.WriteJSON(bet); err != nil {
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
