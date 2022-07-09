package main

import (
	"fmt"
	"math/rand"
	"time"

	"casino.website/pkg/models"
	"casino.website/pkg/services/roulette"
)

func main() {
	// Init rand seed
	rand.Seed(time.Now().UnixNano())

	rGame := roulette.NewRouletteGame()

	go rGame.Start()
	go rGame.End()

	bet1 := &models.RouletteBets{ClientId: "1", BetPosition: "red", BetAmount: 100}
	bet2 := &models.RouletteBets{ClientId: "2", BetPosition: "green", BetAmount: 100}
	bet3 := &models.RouletteBets{ClientId: "2", BetPosition: "black", BetAmount: 100}

	rGame.RegisterBet <- bet1
	rGame.RegisterBet <- bet2
	rGame.RegisterBet <- bet3

	for bet := range rGame.Bets {
		fmt.Printf("Bet : ClientId %s BetPosition %s BetAmount %d\n", bet.ClientId, bet.BetPosition, bet.BetAmount)
	}

	time.Sleep(20 * time.Second)
}
