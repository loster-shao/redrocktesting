package control

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"redrocktest/SQL"
	"redrocktest/token"
	"strconv"
)

func Join(c *gin.Context)  {
	//接受post请求
	tokenValue := c.GetHeader("token")
	num        := c.PostForm("request")
	//password   := c.PostForm("password")
	n, err     := strconv.Atoi(num)
	if err != nil {
		fmt.Println(err)
		c.JSON(502,
			gin.H{
				"status":  http.StatusBadGateway,
				"message": "房间号必须为数字",
			})
	}

	//检查token
	_, username, _ := token.CheckToken(tokenValue)
	fmt.Println(username)

	//连接数据库
	db := SQL.DbConn()
	//r := &room{
	//	Room:     n,
	//	Password: password,
	//}
	//
	//db.Model(&r).Update("player2", username)//名字对应，正在想
	db.Model(Room{}).Where("room=?", n).Update(Room{Player2: username})
}
