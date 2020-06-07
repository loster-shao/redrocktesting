package cookie

import "github.com/gin-gonic/gin"

func Cookie(c *gin.Context, name string)  {
	c.SetCookie(/*1*/"username", name,
		/*2*/6000,
		/*3*/"/",
		/*4*/"localhost",
		/*5*/false,
		/*6*/false)
}
