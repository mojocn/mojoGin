package models

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/mojocn/mojoGin/config"
	"log"
)

var RedisClient *redis.Client
var DB *gorm.DB

func init() {
	//init mysql
	connectionAddr := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	db, err := gorm.Open("mysql", connectionAddr)
	if err != nil {
		log.Fatal(err)
	}
	DB = db

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := RedisClient.Ping().Result()
	log.Println(pong, err)
}
