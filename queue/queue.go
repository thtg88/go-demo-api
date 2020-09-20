package queue

import (
	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/taskq/v3"
	"github.com/vmihailenco/taskq/v3/redisq"
)

var Redis = redis.NewClient(&redis.Options{
	Addr: ":6379",
})

var (
	Factory   = redisq.NewFactory()
	mainQueue = Factory.RegisterQueue(&taskq.QueueOptions{
		Name:  "api-worker",
		Redis: Redis,
	})
)

func AddToMainQueue(task *taskq.Message) error {
	return mainQueue.Add(task)
}
