package oauthcontrollers

import (
	"net/http"

	oauthservices "casino.website/pkg/services/oauth.services"
	"github.com/gin-gonic/gin"
)

func login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if !oauthservices.CheckUserPassword(username, password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Bad credentials"})
		return
	}

	userId := 1

	userClaims := oauthservices.NewUserClaims(userId)

	if userClaims == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured while creating the jwt tokens."})
		return
	}

	ctx.JSON(http.StatusOK, userClaims)
}

func LoginRoutesHandler(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/oauth")

	routerGroup.POST("/login", login)
}
