package main

import (
	"fmt"
	"time"
)

func main() {
	//str1 := "ip saddr 1.1.1.1 ip daddr 1.1.1.1 tcp dport 502 meta time \"2023-11-15 00:00:00\"-\"2023-11-16 00:00:00\" counter packets 0 bytes 0 meta mark set 0x05f5f465 # handle 20"
	////str := "ip saddr 1.1.1.1 ip daddr 1.1.1.1 tcp dport 502 meta time \"2023-11-15 00:00:00\"-\"2023-11-16 00:00:00\" counter packets 0 bytes 10 meta mark set 0x05f5f465 # handle 20"
	//
	//re, _ := regexp.Compile(`bytes \d*`)
	//newStr := re.ReplaceAllString(str1, "bytes 10")
	//fmt.Println(newStr)
	//expireInt := 1727231374
	expireInt := 1702231374
	ExpireDate := time.Unix(int64(expireInt), 0).Format("2006-01-02 15:04:05")
	fmt.Println(ExpireDate)
}
