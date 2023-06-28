package main

import "fmt"

func main() { //并发goroutine、channel、waitGroup
	//c1 := make(chan int) //channel声明,无缓冲区
	//go func() {
	//	c1 <- 1 //没有缓冲区的时候,是没有容量的;只有取的时候是才会有容量(容量在取的时候有),阻塞存值
	//}()
	//fmt.Println(<-c1)
	//
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		c1 <- i //存
	//	}
	//}()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(<-c1) //取,读一次写一次
	//}
	//
	//c2 := make(chan int, 5) //有缓冲区
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		c2 <- i
	//	}
	//}()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(<-c2) //先存满缓冲区,再读一次写一次,类似队列先进先出
	//}
	//
	//ch1 := make(chan int, 5) //在多个goroutine间进行通信
	//var readc <-chan int = ch1
	//var writec chan<- int = ch1
	//writec <- 2
	//fmt.Println(<-readc)
	//fmt.Println()
	//
	//ch1 <- 10
	//ch1 <- 11
	//ch1 <- 12
	//ch1 <- 13
	//ch1 <- 14
	//close(ch1)
	//for v := range ch1 { //遍历ch1中值,必须close,不然读出后会等待输入造成死锁
	//	fmt.Println(v)
	//}
	//
	//ch2 := make(chan int, 1) //select
	//ch3 := make(chan int, 1)
	//ch4 := make(chan int, 1)
	//ch2 <- 2
	//ch3 <- 3
	//ch4 <- 4
	//select {
	//case <-ch2:
	//	fmt.Println("ch2")
	//case <-ch3:
	//	fmt.Println("ch3")
	//case <-ch4:
	//	fmt.Println("ch4") //如果都满足，将随机执行
	//}

	c := make(chan int)
	var read <-chan int = c
	var write chan<- int = c
	go setC(write)
	getC(read)
}

func setC(write chan<- int) {
	for i := 0; i < 10; i++ {
		write <- i
	}
}

func getC(read <-chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-read)
	}
}
