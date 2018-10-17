package data_conn

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type User struct {
	Id       int    `gorm:"auto_increment"`
	Num      string `gorm:"not null"`
	Password string `gorm:"not null"`
}

type QuestionInfo struct {
	Id          int    `gorm:"auto_increment"`
	Question    string `gorm:"not null"`
	Questioner  string `gorm:"not null"`
	AnswerCount int    `gorm:"not null"`
}

type AnswerInfo struct {
	Id       int    `gorm:"not null"`
	Answer   string `gorm:"not null"`
	Answerer string `gorm:"not null"`
}

func MakeMySqlDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root123@(127.0.0.1:3306)/xiangmu?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	// 自动迁移模式
	db.AutoMigrate(&User{}, &QuestionInfo{}, &AnswerInfo{})
	return db
}

func MakeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println(pong, err)
	// Output: PONG <nil>
	return client
}
