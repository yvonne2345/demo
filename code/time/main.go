package main

import (
	"fmt"
	"time"
)

// TaskUpdateSyslogConfig 定时从redis更新内存syslogconfig
func main() {
	// 每20s执行行
	ticker := time.NewTicker(time.Second * 1)
	var i int64
	for range ticker.C {
		i = i + 1
		fmt.Println(i)

	}
}
