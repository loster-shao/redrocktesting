package SQL

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "strconv"
	)

var DB *gorm.DB

func DbConn() *gorm.DB {

	userName:="root"
	password:=""
	host:="127.0.0.1"
	dbName:="test"
	port :=3306
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", userName,password, host, port, dbName )
	db, err := gorm.Open("mysql", connArgs)

	if err != nil {
		fmt.Println("err:", err)
		return nil
	}

	db.SingularTable(true)
	return db
}
