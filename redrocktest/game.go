package main

import (
	"fmt"
	"strconv"
)

var Qi [15][15]int//棋盘大小
var s  int

//游戏
func Game(x, y int) (int, int, int) {
	s++
	fmt.Println("s", s%2)
	Xia(x-1,y-1, s%2)
	Win(x-1, y-1, strconv.Itoa(s%2))
	return x, y, s
}

//棋盘位置
func Xia(x, y, zi int) (int, int) {
	//var m int
	//if s == 0{
	//	m := 1
	//}else {
	//	m := 2
	//}
	//fmt.Println("m:", m)
	Qi[x][y] = zi + 1
	Printn()
	//fmt.Printf("请玩家%d输入位置:", n)
	//Check(x-1, y-1)
	return x, y
}

//打印棋盘
func Printn()  {
	for i:=0; i<15; i++ {
		for j:=0; j<15; j++ {
			fmt.Print(Qi[i][j]," ")
		}
		fmt.Print("\n")
	}
}

//如果有子就重下否则。。。没啥用了这个。。。好多我想的聊天室做不进去
func Check(x, y int) bool {
	if x<16 && y<16 {
		if Qi[x][y] != 0 {
			return false
		}
		return true
	}else {
		fmt.Println("输入不合法，请重新输入：")
	}
	return true
}

//判断胜利
func Win(x, y int, name string) bool {
	n := Qi[x][y]

	//行
	h := 1
	for i:=1; i<5; i++ {
		if x+i < 15 {
			if Qi[x+i][y] == n {
				h = h + 1
			} else {
				break
			}
		}
	}
	for i:=1; i<5; i++ {
		if x-i >= 0 {
			if Qi[x-i][y] == n{
				h +=1
			} else {
				break
			}
		}
	}

	//列
	s := 1
	for i := 1; i < 5; i++ {
		if y+i < 15 {
			if Qi[x][y+i] == n {
				s += 1
			} else {
				break
			}
		}
	}
	for i := 1; i < 5; i++ {
		if y-i >= 0 {
			if Qi[x][y-i] == n{
				s += 1
			} else {
				break
			}
		}
	}

	//右斜
	r := 1
	for i:=1; i<5; i++ {
		if y+i < 15 && x+i < 15 {
			if Qi[x+i][y+i] == n {
				r += 1
			} else {
				break
			}
		}
	}
	for i:=1; i<5; i++{
		if y-i >= 0 && x-i >= 0 {
			if Qi[x-i][y-i] == n{
				r += 1
			} else {
				break
			}
		}
	}

	//左斜
	l := 1
	for i:=1; i<5; i++ {
		if y+i < 15 && x-i >=0 {
			if Qi[x-i][y+i] == n {
				l += 1
			} else {
				break
			}
		}
	}
	for i:=1; i<5; i++{
		if y-i >= 0 && x+i < 15 {
			if Qi[x+i][y-i] == n{
				l +=1
			} else {
				break
			}
		}
	}

	fmt.Println("n:", n, "h:", h, "s:", s, "r:", r, "l:", l)
	if h >= 5 || l >= 5 || r >= 5 || s >= 5 {
		fmt.Printf("恭喜%s胜利", name)
		return true
	}
	return false
}
