package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"time"
)

type delay struct {
	cli    *redis.Client
	expire time.Duration
	psc    *redis.PubSub
}

var client *redis.Client
var psc *redis.PubSub

func Init() {
	viper.SetConfigName("config")
	// 创建 Redis 客户端

	client = redis.NewClient(&redis.Options{
		Addr: viper.GetString("redis.addr"),
		DB:   viper.GetInt("redis.db"),
	})
	// 创建一个订阅上下文
	ctx := context.Background()

	//// 创建订阅对象
	psc = client.Subscribe(ctx, "__keyevent@0__:expired")
}

func NewDelay(expired time.Duration) *delay {
	return &delay{
		cli:    client,
		expire: expired,
		psc:    psc,
	}
}

func (d *delay) Push(key string, value string) error {
	return d.cli.Set(context.Background(), key, value, d.expire).Err()
}

func (d *delay) Pop() *redis.Message {
	ch := d.psc.Channel()
	select {
	case msg := <-ch:
		return msg
	case <-time.After(1 * time.Second):
		return nil
	}
}

func (d *delay) Close() {
	d.psc.Close()
	d.cli.Close()
}
