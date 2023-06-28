package counter

import (
	"sync"
	"sync/atomic"
)

type Counter struct {
	Count int64
}

var instance *Counter
var once sync.Once

// GetCounter singleton 单例模式
func GetCounter() *Counter {
	once.Do(func() { //重点代码，这个只会执行一次
		instance = &Counter{0}
	})
	return instance
}
func (c *Counter) Add() {
	atomic.AddInt64(&c.Count, 1)
}

func (c *Counter) Sub() {
	atomic.AddInt64(&c.Count, -1)
}

func (c *Counter) GetCount() int64 {
	return atomic.LoadInt64(&c.Count)
}

func (c *Counter) Clear() {
	atomic.StoreInt64(&c.Count, 0)
}
