package test

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

var (
	ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
)

func TestSetValue(t *testing.T) {
	err := rdb.Set(ctx, "key", "value", time.Second*10).Err()
	if err != nil {
		t.Fatal(err)
	}
}
func TestGetValue(t *testing.T) {
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(val)
}
