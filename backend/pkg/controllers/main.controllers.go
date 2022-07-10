package controllers

import (
	oauthcontrollers "casino.website/pkg/controllers/oauth.controllers"
	roulettecontrollers "casino.website/pkg/controllers/roulette.controllers"
	"github.com/gin-gonic/gin"
)

func RegisterControllers(rg *gin.RouterGroup) {
	// Register all controllers with the router group
	roulettecontrollers.HandleRouletteGame(rg)
	oauthcontrollers.LoginRoutesHandler(rg)
}
