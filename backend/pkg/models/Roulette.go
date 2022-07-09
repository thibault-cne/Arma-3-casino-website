package models

type RouletteBets struct {
	ClientId    string `json:"id"`
	BetPosition string `json:"position"`
	BetAmount   int    `json:"amount"`
}
