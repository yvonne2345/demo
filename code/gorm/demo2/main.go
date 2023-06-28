package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// 定义模型
type User struct {
	gorm.Model   //内嵌gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"` //结构体标签Tag,可以指定primary key
	Role         string  `gorm:"size:255"`                       // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"`                // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`                 // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`                     // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`                              // 忽略本字段
}

// 使用`AnimalID`作为主键
type Animal struct {
	AnimalID int64  `gorm:"primary_key"`        //自定义主键
	Name     string `gorm:"column:animal_name"` //使用结构体指定列名
	Age      int64  `gorm:"column:animal_age"`  //使用结构体指定列名
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

	db.AutoMigrate(&User{})
	//db.AutoMigrate(&Animal{})
}
