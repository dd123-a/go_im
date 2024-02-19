package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

var DB *gorm.DB

func Database(connString string) {
	db,err:=gorm.Open("mysql",connString)
	if err!=nil{
		fmt.Println("connect err:",err)
	}
	//开启SQL日志模式
	db.LogMode(true)
	if err!=nil{
		//panic(err)
	}
	if gin.Mode()=="release"{
		db.LogMode(false)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(20)//缓存
	db.DB().SetMaxOpenConns(100)//上限
	db.DB().SetConnMaxLifetime(time.Second*30)
	DB=db
	migration()
}














