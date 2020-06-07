package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"redrocktest/control"
)

func SetupRouter(app *gin.Engine)  {
	fmt.Println()
	app.POST("register", control.Register)
	app.POST("login",    control.Login)
	app.POST("logout",   control.Logout)
	app.POST("create",   control.Create)
	app.POST("join",     control.Join)
}
