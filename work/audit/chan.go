package main

import (
	"fmt"
	"sync"
	"time"
)

type Message struct {
	Id  int64
	Msg string
}

var ch = make(chan Message, 200)
var WaitGroup sync.WaitGroup

func main() {
	//var msg Message
	//var msg1 Message
	//msg.Id = 42
	//msg1.Id = 422
	//go func() {
	//	ch <- msg  // 向通道发送一个值
	//	ch <- msg1 // 向通道发送一个值
	//}()
	//go RelationAnalysisTime()
	go Time()
	fmt.Println("====")
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			ticker.Reset(5 * time.Second)
		case value := <-ch:
			fmt.Println(value) // 输出接收到的值
			fmt.Println("====")
			value2 := <-ch      // 从通道中接收值
			fmt.Println(value2) // 输出接收到的值
		}
		//value := <-ch      // 从通道中接收值
		//fmt.Println(value) // 输出接收到的值
		//fmt.Println("====")
		//value2 := <-ch      // 从通道中接收值
		//fmt.Println(value2) // 输出接收到的值
	}
}

func Time() {
	//go func() {
	//
	//}()
	for {
		var msg Message
		msg.Id = 10000054
		ch <- msg
		msg.Id = 24507
		ch <- msg

		fmt.Println(len(ch))
		//fmt.Println("111111")
	}
	time.Sleep(1 * time.Second)
}

//func RelationAnalysisTime() {
//	go func() {
//		for {
//			var msg utils.Message
//			msg.Id = 10000054
//			utils.RelationAnalysisQueue <- msg
//
//			msg.Id = 24507
//			utils.RelationAnalysisQueue <- msg
//			time.Sleep(1 * time.Second)
//		}
//	}()
//}
