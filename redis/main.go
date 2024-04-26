package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var (
	RedisClient *redis.Client
	QueueName   string
)

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "10.25.30.126:6379",
		Password: "Netvine123#@!",
		DB:       1,
	})
	QueueName = "queue"
}

// 推送消息
func PushMessage(message string) error {
	if err := RedisClient.LPush(QueueName, message).Err(); err != nil {
		return err
	}
	return nil
}

// 消费消息
func ConsumeMessage() {
	for {
		result, err := RedisClient.BRPop(0, QueueName).Result()
		if err != nil {
			fmt.Println("consume message error: ", err)
			continue
		}
		fmt.Println("consume message: ", result[1]) // 输出消息内容
	}
}

func main() {
	go ConsumeMessage() // 异步消费消息
	i := 0
	for i < 100 {
		PushMessage(fmt.Sprintf("message:%d", i)) // 推送消息
		time.Sleep(time.Millisecond * 100)
		i++
	}
	select {}
}
