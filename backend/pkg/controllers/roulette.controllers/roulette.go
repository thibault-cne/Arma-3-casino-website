package roulettecontrollers

import (
	"fmt"

	"casino.website/pkg/server/websocket"
	rouletteservices "casino.website/pkg/services/roulette.services"
	"github.com/gin-gonic/gin"
)

func connectRoulette(ctx *gin.Context, rGame *rouletteservices.RouletteGame) {
	// reqToken := ctx.Request.Header.Get("Authorization")
	// claims, err := oauthservices.RetrieveUserClaims(reqToken)

	// if err != nil {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	//	return
	// }

	conn, err := websocket.WsUpgrader(ctx.Writer, ctx.Request)

	if err != nil {
		fmt.Printf("An error occured while upgrading connection : %s.\n", err.Error())
		return
	}

	client := &rouletteservices.RouletteClient{ClientId: "1", WsConn: conn}

	client.ReadRouletteClient(rGame)
}

func HandleRouletteGame(rg *gin.RouterGroup) {
	rGame := rouletteservices.NewRouletteGame()

	routerGroup := rg.Group("/roulette")
	routerGroup.GET("/connect", func(ctx *gin.Context) {
		connectRoulette(ctx, rGame)
	})

	// go rGame.Start()
	// go rGame.End()
}
