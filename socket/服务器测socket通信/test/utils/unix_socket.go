package utils

import (
	"aaa/utils/counter"
	"fmt"
	"net"
	"os"
	"time"
)

const max_data_limit = 50000

type UnixSocket struct {
	filename string
	bufsize  int
	handler  func(string) string
	Queue    chan string
	Counter  *counter.Counter
}

func NewUnixSocket(filename string, size ...int) *UnixSocket {
	size1 := 10480
	if size != nil {
		size1 = size[0]
	}
	queue := make(chan string, 100)
	us := UnixSocket{filename: filename, bufsize: size1, Queue: queue, Counter: counter.GetCounter()}
	return &us
}

func (this *UnixSocket) createServer() {
	fmt.Println("socket监听执行========================================")
	os.Remove(this.filename)
	addr, err := net.ResolveUnixAddr("unixgram", this.filename)
	if err != nil {
		panic("Cannot resolve unix addr: " + err.Error())
	}
	c, err := net.ListenUnixgram("unixgram", addr)
	defer c.Close()
	if err != nil {
		panic("Cannot listen to unix domain socket: " + err.Error())
	}
	os.Chmod(this.filename, 0666)

	// 回调函数时串行的
	go this.handlerQueueDataLoop()

	for {
		data := make([]byte, 4096)
		nr, _, err := c.ReadFrom(data)
		if err != nil {
			fmt.Printf("conn.ReadFrom error: %s\n", err)
			return
		}

		if this.Counter.GetCount() > max_data_limit {
			fmt.Println("数据量已经超过", max_data_limit, "开始丢弃数据 ------------------")
		} else {
			//fmt.Println("数据量", strconv.Itoa(int(this.Counter.GetCount())), " ------------------")
			this.Counter.Add() // 再存入之前计数
			go this.HandleServerConn(c, string(data[0:nr]))
		}
	}

}

//接收连接并处理
func (socket *UnixSocket) HandleServerConn(conn net.Conn, data string) {
	// 数据线送入队列
	// this.HandleServerContext(data)
	socket.Queue <- data
}

func (socket *UnixSocket) handlerQueueDataLoop() {
	for {
		select {
		case data, ok := <-socket.Queue:
			if ok {
				socket.Counter.Sub()
				socket.HandleServerContext(data)
			} else {
				fmt.Println("从socket queue中读取数据失败!")
			}
		}
	}
}

func (this *UnixSocket) SetContextHandler(f func(string) string) {
	this.handler = f
}

//接收内容并返回结果
func (this *UnixSocket) HandleServerContext(context string) string {
	if this.handler != nil {
		return this.handler(context)
	}
	now := time.Now().String()
	return now
}

func (this *UnixSocket) StartServer() {
	this.createServer()
}

//客户端
func (this *UnixSocket) ClientSendContext(context string) {
	addr, err := net.ResolveUnixAddr("unixgram", this.filename)
	if err != nil {
		panic("Cannot resolve unix addr: " + err.Error())
	}
	//拔号
	c, err := net.DialUnix("unixgram", nil, addr)
	if err != nil {
		panic("DialUnix failed.")
	}
	//写出
	_, err = c.Write([]byte(context))
	if err != nil {
		panic("Writes failed.")
	}
}
