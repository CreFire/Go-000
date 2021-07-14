package main

import (
	"context"
	"fmt"
	goredis "github.com/go-redis/redis/v8"
)

var(
	redisDB         *goredis.Client    // 游戏redis连接客户端
)

func main() {
	ctx := context.Background()
	redisDB = initRedis()
	key := fmt.Sprintf("test_sort")
	for i := 1;i<100;i++{
		redisDB.ZAdd(ctx,key,&goredis.Z{Score:float64(i),Member:fmt.Sprintf("redis_%v",i)})
	}
	redisDB.ZAdd(ctx,key,&goredis.Z{Score: 1.23,Member:fmt.Sprintf("redis_%v",1.23)})
	strings, err := redisDB.ZRangeWithScores(ctx, key, 0, 20).Result()
	if err != nil {
		fmt.Println(err)
	}
	for key, value := range strings {
		fmt.Println(key,":",value.Score,":",value.Member)
	}
}


func initRedis() *goredis.Client {
	return goredis.NewClient(&goredis.Options{
		Addr:        "10.255.0.10:6380",
		Password:    "",
		DB:          10,
		IdleTimeout: 300,
		DialTimeout: 10,
	})
}