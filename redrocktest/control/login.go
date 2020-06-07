package control

import (
	"fmt"
	"redrocktest/SQL"
)


//登录
func Login(username, password string/*c *gin.Context*/){
	//username := c.PostForm("username")
	//password := c.PostForm("password")
	db       := SQL.DbConn()//函数不清楚。。。
	defer db.Close()

	var u user
	db.Where("username = ?", username).First(&u)
	//tokens := token.Create(u.Username, int(u.Model.ID))//创建token
	if u.Username == ""{
		fmt.Println("no")
		Register(username, password)
	}
	if password != u.Password {
		//v.Write
	//	c.JSON(200,gin.H{"status:": http.StatusOK, "message":"登录成功", "token": &tokens})
	//}else {
	//	c.JSON(400,gin.H{"status:": http.StatusBadRequest, "message": "用户名或密码错误"})
	}
}