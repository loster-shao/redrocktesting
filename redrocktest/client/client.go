package main

import (
	"fmt"
	"net"
)

var username string = "" //声明聊天室的昵称
var password string = ""

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080") //打开监听端口
	if err != nil {
		fmt.Println("conn fail...")
	}
	defer conn.Close()
	fmt.Println("client connect server successed \n")

	//给自己取一个聊天室的昵称
	fmt.Printf("请输入你的账号或密码:")
	fmt.Scanf("%s%s", &username, &password)             //输入昵称
	fmt.Println("hello : ", username)      //客户端输出
	conn.Write([]byte("nick|" + username + "|" + password)) //将信息发送给服务器端

	go Handle(conn) //创建线程

	var x, y string
	for {
		fmt.Scan(&x, &y)                                //输入消息
		conn.Write([]byte("say|" + username + "|" + x + "|" + y)) //三段字节流 say | 昵称 | 发送的消息
		//if x == "" || y == "" {                            //如果消息为quit
		//	conn.Write([]byte("quit|" + username)) //将quit字节流发送给服务器端
		//	break                              //程序结束运行
		//}
	}
}

func Handle(conn net.Conn) {

	for {

		data := make([]byte, 255)        //创建一个字节流
		msg_read, err := conn.Read(data) //将读取的字节流赋值给msg_read和err
		if msg_read == 0 || err != nil { //如果字节流为0或者有错误
			break
		}
		fmt.Println(string(data[0:msg_read])) //把字节流转换成字符串
	}
}
