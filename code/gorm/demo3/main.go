package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 1、定义模型
type User struct { //根据变量定义的结构体确定查哪张表
	gorm.Model
	Name string `gorm:"default:'ABC'"`
	Age  int
}

func main() {
	//连接数据库
	dsn := "root:123456@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{ //修改默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	//2、把模型和数据库的表对应
	db.AutoMigrate(&User{})

	//3.初始化结构体-赋值
	//u1 := User{Name: "bill", Age: 26}
	//db.Create(&u1)
	//u2 := User{Name: "clavin", Age: 27}
	//db.Create(&u2)

	//4、查询
	//var user User //声明模型结构体类型变量user
	user := new(User) //make给channel、map、slice申请内存初始化操作，类型不改变;new基本数据类型、结构体，返回对应类型的指针
	//立即执行方法
	//db.Debug().First(user) //第一条记录 Debug()打印sql语句 SELECT * FROM `user` WHERE `user`.`deleted_at` IS NULL ORDER BY `user`.`id` LIMIT 1
	//db.Debug().Take(user) //随机记录 SELECT * FROM `user` WHERE `user`.`deleted_at` IS NULL LIMIT 1
	//db.Debug().Last(user) //最后一条记录 SELECT * FROM `user` WHERE `user`.`deleted_at` IS NULL ORDER BY `user`.`id` DESC LIMIT 1
	var users []User //声明接收所有表结果的结构体切片users—数组
	//db.Debug().Find(&users) //所有记录 SELECT * FROM `user` WHERE `user`.`deleted_at` IS NULL
	//db.Debug().Find(user) //一条记录,因为是user一个结构体实例,users是切片 SELECT * FROM `user` WHERE `user`.`deleted_at` IS NULL

	//where
	//db.Debug().Where("name=?", "mary").First(user) //SELECT * FROM `user` WHERE name='mary' AND `user`.`deleted_at` IS NULL ORDER BY `user`.`id` LIMIT 1
	//db.Debug().Where("name=?", "mary").Find(&users) //SELECT * FROM `user` WHERE name='mary' AND `user`.`deleted_at` IS NULL
	//db.Debug().Where("name<>?", "mary").Find(&users) //非 SELECT * FROM `user` WHERE name<>'mary' AND `user`.`deleted_at` IS NULL
	//db.Debug().Where("name in (?)", []string{"mary", "ben"}).Find(&users) //in SELECT * FROM `user` WHERE name in ('mary','ben') AND `user`.`deleted_at` IS NULL
	//db.Debug().Where("name like ?", "%e%").Find(&users) //like SELECT * FROM `user` WHERE name like '%e%' AND `user`.`deleted_at` IS NULL
	//db.Debug().Where("name=? and age>=?", "mary", "20").Find(&users) //SELECT * FROM `user` WHERE (name='mary' and age>='20') AND `user`.`deleted_at` IS NULL

	//更新update-单个属性、updates-多个

	//save-更新所有字段

	fmt.Printf("user:%#v\n", users)
	fmt.Printf("user:%#v\n", user)
}
