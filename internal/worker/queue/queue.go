package queue

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/taskq/v3"
	"github.com/vmihailenco/taskq/v3/redisq"
)

// Redis is the Redis client
var Redis *redis.Client

// Factory is the taskq Factory
var Factory taskq.Factory

// MainQueue is the main queue to put tasks on
var MainQueue taskq.Queue

// AddToMainQueue adds a given taskq.Message to the MainQueue
func AddToMainQueue(task *taskq.Message) error {
	return MainQueue.Add(task)
}

// InitializeRedis initalizes a Redis connection from the os environment variables
func InitializeRedis() {
	opts := GetRedisOptions()
	Redis = redis.NewClient(opts)
	Factory = redisq.NewFactory()
	MainQueue = Factory.RegisterQueue(&taskq.QueueOptions{
		Name:  "api-worker",
		Redis: Redis,
	})
}

// GetRedisOptions returns the options for the Redis connections
// from the environment variables
func GetRedisOptions() *redis.Options {
	var address string
	var password string
	var username string
	var redisDb int

	urlStr := os.Getenv("REDIS_URL")

	if urlStr != "" {
		u, err := url.Parse(urlStr)
		if err != nil {
			panic(err)
		}

		address = u.Host
		userInfo := strings.Split(u.User.String(), ":")
		username = userInfo[0]
		if len(userInfo) > 1 {
			password = userInfo[1]
		}
		redisDb, err = strconv.Atoi(u.Path)
		if err != nil {
			redisDb = 0
		}
	} else {
		address = fmt.Sprintf(
			"%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
		)
		password = os.Getenv("REDIS_PASSWORD")
		redisDb, _ = strconv.Atoi(os.Getenv("REDIS_DATABASE"))
		username = os.Getenv("REDIS_USERNAME")
	}

	opts := &redis.Options{
		Addr:     address,
		Password: password,
		DB:       redisDb,
	}

	// if "h" is provided as username we assume that
	// the Redis instance does not support providing the username e.g. < v6
	if username != "h" {
		opts.Username = username
	}

	return opts
}
