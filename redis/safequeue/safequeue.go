package main

// A simple demo implementation of a Redis job queue in Go inspired on http://stackoverflow.com/a/34754632
// You'll need to get redis driver package in terminal with: go get -u gopkg.in/redis.v5
// Once it is running, Redis should look like: http://i.imgur.com/P4XlwlP.png
// Terminal should look like: http://i.imgur.com/AS2IIbP.png
// I have 4 days of Go programming experience, have mercy.

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	redis "gopkg.in/redis.v5"
)

func main() {
	// connect to Redis
	redisParams := redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0}
	client := redis.NewClient(&redisParams)
	if _, err := client.Ping().Result(); err != nil {
		panic("Could not ping Redis: " + err.Error())
	}

	// cleanup Redis from any residual mess of previous runs
	client.Del("queue:tasks")
	client.Del("queue:processing")
	client.Del("queue:processing:lock")

	// fire up 3 consumer goroutines (aka "threads")
	for i := 0; i < 3; i++ {
		go consumetasks("Worker#"+strconv.Itoa(i+1), redisParams)
	}
	// fire up 1 rescheduledeadtasks()
	go rescheduledeadtasks(redisParams)

	// produce 10 tasks every 10 seconds to be consumed by consumer goroutines we fired just above
	taskIDStart := 1
	for {
		produce10tasks(redisParams, taskIDStart)
		taskIDStart += 10
		time.Sleep(10 * time.Second)
	}
}

func rescheduledeadtasks(redisParams redis.Options) {
	client := redis.NewClient(&redisParams)
	if _, err := client.Ping().Result(); err != nil {
		panic("Could not ping Redis: " + err.Error())
	}

	// loop endlessly
	for {
		// sleep 2 seconds because we don't need to waste CPU cycles checking for
		// dead workers thousands of time per second
		time.Sleep(2 * time.Second)

		// get all tasks being processed
		tasks, err := client.LRange("queue:processing", 0, -1).Result()
		if err != nil {
			panic("Error while fetching queue:processing" + err.Error())
		}

		// foreach job in jobs : lock = GET lock:processing-q:job_identifier
		// if lock is null this job timed out, retry with RPUSH q "job_identifier"
		for _, task := range tasks {
			lockKey := "queue:processing:lock:" + task
			_, err := client.Get(lockKey).Result()
			if err == redis.Nil {
				fmt.Println(task + " timed out. Reschedulling.")
				if err := client.RPush("queue:tasks", task).Err(); err != nil {
					panic("Could not reschedule task to list: " + err.Error())
				}
			} else if err != nil {
				panic("Error while fetching " + lockKey + ": " + err.Error())
			}
		}

	}
}

func produce10tasks(redisParams redis.Options, taskIDStart int) {
	client := redis.NewClient(&redisParams)
	if _, err := client.Ping().Result(); err != nil {
		panic("Could not ping Redis: " + err.Error())
	}

	// add 100 items to queue
	for taskID := 0; taskID < 10; taskID++ {
		if err := client.RPush("queue:tasks", "task"+strconv.Itoa(taskIDStart+taskID)).Err(); err != nil {
			panic("Could not add task to list: " + err.Error())
		}
	}
}

func consumetasks(clientID string, redisParams redis.Options) {
	client := redis.NewClient(&redisParams)
	if _, err := client.Ping().Result(); err != nil {
		panic(clientID + ": Could not ping Redis: " + err.Error())
	}

	// loop endlessly
	for {
		//fmt.Println(clientID + ": Waiting for new tasks")
		// fetch 1 task, move it to "processing"" queue
		taskID, err := client.BRPopLPush("queue:tasks", "queue:processing", 0).Result()
		if err != nil {
			panic(clientID + ": Error while reading task from queue: " + err.Error())
		}

		// create a key that will delete itself in 60 seconds. the cleanuptasks() checks if there's any
		// key in processing queue that isn't on lock queue to detect dead workers
		client.Set("queue:processing:lock:"+taskID, "1", 2*time.Second)

		//fmt.Println(clientID + ": Processing task " + taskID)

		// sleep between 1 and 3 seconds
		sleepTime := time.Duration(1 + rand.Intn(3))
		time.Sleep(sleepTime * time.Second)

		fmt.Println(clientID + ": Done task " + taskID)
		client.LRem("queue:processing", 0, taskID)
	}
}
