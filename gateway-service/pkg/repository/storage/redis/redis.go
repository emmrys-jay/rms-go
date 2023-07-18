package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"gateway-service/internal/config"
	"log"
	"time"

	"gateway-service/utility"
	"github.com/go-redis/redis/v8"
)

var (
	Rds *redis.Client
	Ctx = context.Background()
)

func SetupRedis() {
	logger := utility.NewLogger()
	getConfig := config.GetConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", getConfig.Redis.Host, getConfig.Redis.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := rdb.Ping(Ctx).Err(); err != nil {
		fmt.Printf("%v:%v", getConfig.Redis.Host, getConfig.Redis.Port)
		log.Fatalln("Redis db error: ", err)
	}
	pong, _ := rdb.Ping(Ctx).Result()
	fmt.Println("Redis says: ", pong)
	Rds = rdb
	logger.Info("Redis CONNECTION ESTABLISHED")
}

type Redis struct {
	Rdb *redis.Client
}

func GetRedisDb() RedisRepository {
	return &Redis{Rdb: Rds}
}

func (rdb *Redis) RedisSet(key string, value interface{}) error {
	serialized, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return rdb.Rdb.Set(Ctx, key, serialized, 24*time.Hour).Err()
}

func (rdb *Redis) RedisGet(key string) ([]byte, error) {
	serialized, err := rdb.Rdb.Get(Ctx, key).Bytes()
	return serialized, err
}

func (rdb *Redis) RedisDelete(key string) (int64, error) {
	deleted, err := rdb.Rdb.Del(Ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
