package storage

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

func RedisTest() {
	// "redis_conf": {
	//     "host": "redis://redis:d9IU3Wwu06@itihas-redis-master.de-stage-itihas",
	//     "port": 6379,
	//     "db": 0,
	//     "bucket_size": 10000,
	//     "ttl_day_end": false,
	//     "ttl": "2m"
	//   },
	rdb := redis.NewClient(&redis.Options{
		Addr:     "itihas-redis-master:6379",
		Password: "d9IU3Wwu06", // no password set
		DB:       0,            // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
