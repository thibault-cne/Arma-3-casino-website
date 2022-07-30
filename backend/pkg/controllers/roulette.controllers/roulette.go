package roulettecontrollers

import (
	"fmt"
	"strconv"

	"casino.website/pkg/config/server/websocket"
	oauthservices "casino.website/pkg/services/oauth.services"
	rouletteservices "casino.website/pkg/services/roulette.services"
	"github.com/gin-gonic/gin"
)

func connectRoulette(ctx *gin.Context, rGame *rouletteservices.RouletteGame) {
	conn, err := websocket.WsUpgrader(ctx.Writer, ctx.Request)

	if err != nil {
		fmt.Printf("An error occured while upgrading connection : %s.\n", err.Error())
		return
	}

	reqToken := ctx.Query("token")
	claims, err := oauthservices.DecodeToken(reqToken)

	if err != nil {
		rGame.RegisterPlayerConn <- conn

	} else {
		client := &rouletteservices.RouletteClient{ClientId: strconv.Itoa(claims.User_id), WsConn: conn}

		rGame.RegisterPlayerConn <- client.WsConn

		client.ReadRouletteClient(rGame)
	}
}

func HandleRouletteGame(rg *gin.RouterGroup) {
	rGame := rouletteservices.NewRouletteGame()

	routerGroup := rg.Group("/roulette")
	routerGroup.GET("/connect", func(ctx *gin.Context) {
		connectRoulette(ctx, rGame)
	})

	go rGame.Start()
	// go rGame.End()
}
