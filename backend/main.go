package main

import (
	"math/rand"
	"time"

	"casino.website/pkg/db"
	"casino.website/pkg/server"
)

func main() {
	// Init rand seed
	rand.Seed(time.Now().UnixNano())

	// Init the database
	db.InitDatabase()

	// Init the htpp server
	server.InitServer()
}
