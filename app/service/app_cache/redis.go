package app_cache

import (
	"errors"
	"goskeleton/app/utils/redis_factory"
	"log"
)

type Redis struct {
	client *redis_factory.RedisClient
}

// RedisClient 获取一个redis客户端
func RedisClient() *Redis {
	return &Redis{client: redis_factory.GetOneRedisClient()}
}

// closeClient 释放客户端
func (r *Redis) closeClient() {
	r.client.ReleaseOneRedisClient()
}

func (r *Redis) Set(key string, val interface{}) error {
	defer r.closeClient()

	_, err := r.client.Execute("set", key, val)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *Redis) SetEX(key string, val interface{}, ttl int) error {
	defer r.closeClient()

	_, err := r.client.Execute("setex", key, ttl, val)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *Redis) SetNX(key string, val interface{}) error {
	defer r.closeClient()

	_, err := r.client.Execute("setnx", key, val)
	if err != nil {
		return err
	}

	return nil
}

func (r *Redis) Get(key string) (string, error) {
	defer r.closeClient()

	if res, err := r.client.String(r.client.Execute("get", key)); err != nil {
		return "", err
	} else {
		return res, nil
	}
}

// Delete 删除某个键
func (r *Redis) Delete(key string) error {
	defer r.closeClient()

	if _, err := r.client.Execute("del", key); err != nil {
		return err
	} else {
		return nil
	}
}

// Flush redis不可以清空数据库
func (r *Redis) Flush() error {
	return errors.New("请勿清空 Redis 缓存")
}
