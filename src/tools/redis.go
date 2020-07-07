package tools

import (
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

type RedisClient struct {
	Conn *redis.Client
}

var client *redis.Client

func InitRedis(config *TcpConfig) {

	client = redis.NewClient(&redis.Options{
		Addr:               config.Redis.Host + ":" + strconv.Itoa(config.Redis.Port),
		Password:           config.Redis.Password,
		DB:                 config.Redis.Db,
		PoolSize:           config.Redis.PoolSize,                                 // Redis连接池大小
		MaxRetries:         config.Redis.MaxRetries,                               // 最大重试次数
		IdleTimeout:        time.Duration(config.Redis.IdleTimeout) * time.Second, // 空闲链接超时时间
		IdleCheckFrequency: time.Duration(config.Redis.IdleCheckFrequency),
	})

}
func GetRedisClient() *redis.Client {
	return client
}

func (client *RedisClient) Set(key string, value interface{}, expiration time.Duration) *redis.Client {
	err := client.Conn.Set(key, value, expiration).Err()
	if err != nil {
		panic(err)
	}
	return (*client).Conn
}

func (client *RedisClient) Incr(key string) int64 {
	res, err := client.Conn.Incr(key).Result()
	if err != nil {
		panic(err)
	}
	return res
}

func (client *RedisClient) Pipeline() redis.Pipeliner {
	return client.Conn.Pipeline()
}

func (client *RedisClient) Decr(key string) int64 {
	res, err := client.Conn.Decr(key).Result()
	if err != nil {
		panic(err)
	}
	return res
}

func (client *RedisClient) DecrBy(key string, de int64) int64 {
	res, err := client.Conn.DecrBy(key, de).Result()
	if err != nil {
		panic(err)
	}
	return res
}

func (client *RedisClient) Expire(key string, expiration time.Duration) bool {
	res, err := client.Conn.Expire(key, expiration).Result()
	if err != nil {
		panic(err)
	}
	return res
}

func (client *RedisClient) Get(key string) (string, error) {
	val, err := (*client).Conn.Get(key).Result()

	if err == redis.Nil {
		return "", nil
	}

	if err != nil {
		return "", err
	}

	return val, nil
}

func (client *RedisClient) IsExist(key string) bool {
	_, err := (*client).Conn.Get(key).Result()
	return err != redis.Nil
}

func (client *RedisClient) Lpop(key string) (string, error) {
	val, err := (*client).Conn.LPop(key).Result()

	if err == redis.Nil {
		return "", nil
	}

	if err != nil {
		return "", err
	}

	return val, nil
}

func (client *RedisClient) Lpush(key string, values ...interface{}) bool {
	_, err := (*client).Conn.LPush(key, values...).Result()

	return err == nil
}

func (client *RedisClient) Lrange(key string, start, stop int64) ([]string, error) {
	res, err := (*client).Conn.LRange(key, start, stop).Result()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (client *RedisClient) Del(key ...string) *redis.Client {
	_, err := client.Conn.Del(key...).Result()
	if err != nil {
		panic(err)
	}
	return (*client).Conn
}

func (client *RedisClient) SetIfNotExist(key string, value interface{}, expiration time.Duration) bool {
	result, err := (*client).Conn.SetNX(key, value, expiration).Result()
	if err != nil {
		return false
	}
	return result
}

func (client *RedisClient) PSubscribe(channels ...string) *redis.PubSub {
	ps := (*client).Conn.PSubscribe(channels...)
	return ps
}
