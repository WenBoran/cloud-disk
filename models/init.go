package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"log"
	"xorm.io/xorm"
)

var Engine = Init()
var RDB = InitRedis()

func Init() *xorm.Engine {

	engine, err := xorm.NewEngine("mysql", "root:123@/cloud-disk?charset=utf8")
	if err != nil {
		log.Printf("Xorm New Engine Error:%v", err)
	}
	return engine
}

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
