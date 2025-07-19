package db

import (
	"context"
	"dify_feishu_mcp/logger"
	"dify_feishu_mcp/model"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisManager struct {
	Config *model.Config
	client *redis.Client
}

func NewRedisManager(config *model.Config) *RedisManager {
	return &RedisManager{
		Config: config,
	}
}

func (rm *RedisManager) NewRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     rm.Config.RedisAddr,
		Password: rm.Config.RedisPass,
		DB:       rm.Config.RedisDB,
	})
	ctx := context.Background()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		logger.Logger.Panic(err)
	}
	logger.Logger.Info("redis连接成功")
	return rdb, nil
}
func (rm *RedisManager) StopRedis() error {
	if rm.client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := rm.client.Shutdown(ctx).Err(); err != nil {
			logger.Logger.Errorf("Redis关闭异常: %v", err)
		} else {
			logger.Logger.Info("Redis连接已关闭")
		}
		rm.client = nil
	}
	return nil
}

func (rm *RedisManager) GetRedisClient() *redis.Client {
	return rm.client
}
