package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"redrocktest/router"
	)

func main()  {
	fmt.Println("System start!")
	app := gin.Default()

	router.SetupRouter(app)

	app.Run(":8080")
}
