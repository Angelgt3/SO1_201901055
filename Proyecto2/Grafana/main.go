package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	rhost := os.Getenv("REDIS_HOST")
	rauth := os.Getenv("REDIS_AUTH")

	rdb := redis.NewClient(&redis.Options{
		Addr:     rhost + ":6379",
		Password: rauth,
		DB:       0,
	})

	for {
		team := rand.Intn(2) + 1
		err := rdb.HIncrBy(ctx, "teams", "team"+strconv.Itoa(team), 1).Err()
		if err != nil {
			panic(err)
		}
		err = rdb.Incr(ctx, "total").Err()
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(os.Stderr, "%v\n", map[string]int{"team" + strconv.Itoa(team): 1})
		time.Sleep(300 * time.Millisecond)
	}
}
