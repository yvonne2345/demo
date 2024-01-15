package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

func conn() {
	url := "root:Netvine123#@!@tcp(10.25.30.131:3306)/audit?charset=utf8&parseTime=True&loc=Local"
	//连接数据库,连接数据库时，可以加上一些高级配置,就是gorm.Config中的参数
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	db1, _ := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		fmt.Println("连接失败")
		return
	} else {
		fmt.Println("连接成功")
	}
	var pSid []int
	var pStr string
	db.Table("policy_intrusion_dictionary").Select("sid").Find(&pSid)

	for _, v := range pSid {
		pStr = pStr + "," + strconv.Itoa(v)
	}
	var bSid []int
	var bStr string
	db1.Table("security_blacklist_vulnerability").Select("sid").Find(&bSid)
	for _, v := range bSid {
		bStr = bStr + "," + strconv.Itoa(v)
	}
	var sameSid []int
	for _, b := range bSid {
		for _, p := range pSid {
			if b == p {
				sameSid = append(sameSid, b)
			}
		}
	}

	fmt.Println(sameSid)
	fmt.Println(len(sameSid))
}
