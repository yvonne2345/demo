package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	//连接数据库
	dsn := "root:123456@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/db1?charset=utf8mb4&pasreTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	//创建表
	db.AutoMigrate(&UserInfo{})

	//创建数据行
	//u1 := UserInfo{1, "a", "man", "游泳"}
	//db.Create(&u1)

	//查询
	var u UserInfo
	db.First(&u) //第一条

	//更新
	db.Model(&u).Update("hobby", "运动")
	fmt.Printf("u:%#v\n", u)

	//删除
	db.Delete(&u)
}
