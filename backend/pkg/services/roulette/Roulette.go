package roulette

import (
	"fmt"
	"math/rand"
	"time"

	"casino.website/pkg/models"
	"casino.website/pkg/services"
)

type RouletteGame struct {
	Id                   string
	RegisterBet          chan *models.RouletteBets
	RegisterPlayerConn   chan *models.RouletteClient
	UnregisterPlayerConn chan *models.RouletteClient
	Bets                 map[*models.RouletteBets]bool
	ConnectedPlayers     map[*models.RouletteClient]bool
}

func NewRouletteGame() *RouletteGame {
	return &RouletteGame{
		Id:                   services.GenerateRandomId(),
		RegisterBet:          make(chan *models.RouletteBets),
		RegisterPlayerConn:   make(chan *models.RouletteClient),
		UnregisterPlayerConn: make(chan *models.RouletteClient),
		Bets:                 make(map[*models.RouletteBets]bool),
		ConnectedPlayers:     make(map[*models.RouletteClient]bool),
	}
}

func (rGame *RouletteGame) Start() {
	for {
		select {
		case bet := <-rGame.RegisterBet:
			rGame.Bets[bet] = true
		case client := <-rGame.RegisterPlayerConn:
			rGame.ConnectedPlayers[client] = true
		case client := <-rGame.UnregisterPlayerConn:
			delete(rGame.ConnectedPlayers, client)
		}
	}
}

func (rGame *RouletteGame) End() {
	time.Sleep(5 * time.Second)

	result := rand.Intn(3)
	color := ""
	switch result {
	case 0:
		color = "black"
	case 1:
		color = "red"
	case 2:
		color = "green"
	}

	fmt.Printf("The good color is : %s\n", color)

	for bet := range rGame.Bets {
		if bet.BetPosition == color {
			fmt.Printf("Client %s won %d\n", bet.ClientId, bet.BetAmount)
		}
	}
}
