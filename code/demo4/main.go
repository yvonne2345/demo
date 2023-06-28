package main

import (
	"fmt"
	"sync"
	"time"
)

func main() { //sync包,Mutex,RWMutex,Once,WaitGroup,sync.Map
	l := &sync.Mutex{} //用指针,确保用同一把 互斥锁
	go lockFun(l)
	go lockFun(l)
	go lockFun(l)
	time.Sleep(4 * time.Second)

	r := &sync.RWMutex{} //读写锁,读锁和写锁分离开
	go ReadLockFun(r)
	go ReadLockFun(r)

	o := &sync.Once{} // 只执行一次
	for i := 0; i < 10; i++ {
		o.Do(func() {
			fmt.Println(i)
		})
	}

	wg := &sync.WaitGroup{} //并发等待组
	wg.Add(2)               //添加等待数量,负数表示任务减1
	go func() {
		time.Sleep(8 * time.Second)
		wg.Done() //等待数量减1
		fmt.Println("1")
	}()
	go func() {
		time.Sleep(6 * time.Second)
		wg.Done() //配套Add使用
		fmt.Println("2")
	}()
	wg.Wait() //使goroutine在此等待

	m := &sync.Map{} //并发安全字典,原生map在多个goroutine同时往Map中添加数据时，会导致部分添加数据的丢失
	go func() {
		for {
			m.Store(1, 1) //设置key-value值
			m.Store(2, 2)
			m.LoadOrStore(3, 3) //如果key存在则返回key对应的value,否则设置
		}
	}()
	go func() {
		for {
			fmt.Println(m.Load(2)) //根据key获取value
		}
	}()
	time.Sleep(100) //sleep要在主程序中

	m1 := &sync.Map{}
	m1.Store(1, 1) //设值
	m1.Store(2, 2)
	m1.Store(3, 3)
	m1.Delete(3) //直接删除key
	m1.Range(func(key, value interface{}) bool { //遍历 无序
		fmt.Println(key, value)
		time.Sleep(1 * time.Second)
		return true
	})
}

func lockFun(lock *sync.Mutex) { //互斥锁,保证每次只有一个goroutine访问同步代码块中的资源
	lock.Lock()
	fmt.Println("a")
	time.Sleep(1 * time.Second)
	lock.Unlock()
}

func ReadLockFun(lock *sync.RWMutex) { //读写锁
	lock.RLock() //读取的时候不阻塞其他读锁，排斥写锁
	fmt.Println("b")
	time.Sleep(1 * time.Second)
	lock.RUnlock()
}
