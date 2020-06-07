package control

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"redrocktest/SQL"
	"redrocktest/token"
	"strconv"
)

type Room struct {
	Room     int
	Password string
	Player1  string
	Player2  string
}

func Create(c *gin.Context)  {
	//post接受数据
	tokenValue := c.GetHeader("token")
	room_num   := c.PostForm("room")
	password   := c.PostForm("password")
	pw, err    := strconv.Atoi(room_num)
	if err != nil{
		fmt.Println("err!=", err)
		c.JSON(502,
			gin.H{
				"status":  http.StatusBadGateway,
				"message": "房间号必须为数字",
		})
	}

	//检查token
	_, username, _ := token.CheckToken(tokenValue)
	fmt.Println(username)

	//数据库连接
	db := SQL.DbConn()
	defer db.Close()

	//数据库操作
	var r []Room
	db.Where("room=?", room_num).Find(&r)
	if len(r) != 0{
		c.JSON(502,
			gin.H{
				"status": http.StatusBadGateway,
				"message": "房间已存在，请重新修改",
			})
		return
	}

	//注册
	db.AutoMigrate(&Room{})
	db.Create(&Room{
		Room:     pw,
		Password: password,
		Player1:  username,                             //读取姓名有问题，在考虑用cookie还是token
	})
	c.JSON(200,
		gin.H{
			"status":  http.StatusOK,
			"message": "创建房间成功",
			"您的房间号为":  room_num,
		})
}

