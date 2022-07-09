package roulettecontrollers

import (
	"fmt"
	"net/http"

	"casino.website/pkg/server"
	claimsservices "casino.website/pkg/services/claims.services"
	rouletteservices "casino.website/pkg/services/roulette.services"
	"github.com/gin-gonic/gin"
)

func connectRoulette(ctx *gin.Context, rGame *rouletteservices.RouletteGame) {
	reqToken := ctx.Request.Header.Get("Authorization")
	claims, err := claimsservices.RetrieveUserClaims(reqToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	conn, err := server.WsUpgrader(ctx.Writer, ctx.Request)

	if err != nil {
		fmt.Printf("An error occured while upgrading connection : %s.\n", err.Error())
		return
	}

	client := &rouletteservices.RouletteClient{ClientId: claims.Id, WsConn: conn}

	client.ReadRouletteClient(rGame)
}

func HandleRouletteGame(rg *gin.RouterGroup) {
	rGame := rouletteservices.NewRouletteGame()

	routerGroup := rg.Group("/roulette")
	routerGroup.GET("/connect", func(ctx *gin.Context) {
		connectRoulette(ctx, rGame)
	})

	go rGame.Start()
	go rGame.End()
}
