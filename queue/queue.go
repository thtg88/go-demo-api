package queue

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/taskq/v3"
	"github.com/vmihailenco/taskq/v3/redisq"
)

var Redis *redis.Client
var Factory taskq.Factory
var MainQueue taskq.Queue

func AddToMainQueue(task *taskq.Message) error {
	return MainQueue.Add(task)
}

func InitializeRedis() {
	redisDb, _ := strconv.Atoi(os.Getenv("REDIS_DATABASE"))

	Redis = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(
			"%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
		),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisDb,
	})

	Factory = redisq.NewFactory()

	MainQueue = Factory.RegisterQueue(&taskq.QueueOptions{
		Name:  "api-worker",
		Redis: Redis,
	})
}
