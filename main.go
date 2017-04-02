package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

var client *redis.Client

const redisAddr = "localhost:6379"

func redisOptions() *redis.Options {
	return &redis.Options{
		Addr:               redisAddr,
		DB:                 15,
		DialTimeout:        10 * time.Second,
		ReadTimeout:        30 * time.Second,
		WriteTimeout:       30 * time.Second,
		PoolSize:           10,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        500 * time.Millisecond,
		IdleCheckFrequency: 500 * time.Millisecond,
	}
}

func init() {
	client = redis.NewClient(redisOptions())
}

func main() {
	pubsub, err := client.Subscribe("mychannel")
	if err != nil {
		log.Fatal(err)
	}
	defer pubsub.Close()
	// launch go routine to listen msgs
	go func() {
		for {
			msg, err := pubsub.ReceiveMessage()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(msg)
		}
	}()

	client.Publish("mychannel", "ola")
	client.Publish("mychannel", "hello")
	pubsub.Unsubscribe("mychannel")
}
