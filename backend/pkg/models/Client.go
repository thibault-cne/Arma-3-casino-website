package models

type Client struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Wallet   int    `json:"wallet"`
}
