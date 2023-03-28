package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB  *DataBase
	Rdb *redis.Client
)

type DataBase struct {
	*gorm.DB
}

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app initialized")
}

func InitMySQL() {
	//自定义日志模板，打印sql语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			Colorful:      true,        //彩色
			LogLevel:      logger.Info, //级别
		},
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.user"), viper.GetString("mysql.password"),
		viper.GetString("mysql.host"), viper.GetString("mysql.port"), viper.GetString("mysql.db_name"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	fmt.Println("config MySQL initialized")
	//db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	database := new(DataBase)
	database.DB = db
	DB = database
}

func InitRedis() {
	//r := gin.Default()
	dsn := fmt.Sprintf("%s:%s", viper.GetString("redis.addr"), viper.GetString("redis.port"))
	Rdb = redis.NewClient(&redis.Options{
		Addr:         dsn,
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.db"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	pong, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("init redis error:", err)
	} else {
		fmt.Printf("redis ping result: %s\n", pong)
	}
}

const (
	PublishKey = "websocket"
)

// Publish 发布消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	fmt.Println("publish", msg)
	err := Rdb.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println("Publish err", err)
	}
	return err
}

// Subscribe 订阅Redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Rdb.Subscribe(ctx, channel)
	fmt.Println("subscribe ctx1", ctx)
	message, err := sub.ReceiveMessage(ctx)
	//fmt.Println("subscribe ctx2", ctx)
	if err != nil {
		fmt.Println("sub.ReceiveMessage err", err)
	}
	fmt.Println("subscribe message:", message)
	return message.Payload, err
}
