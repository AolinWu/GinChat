package utils

import (
	"context"
	"errors"
	"ginchat/config"
	"ginchat/log"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Client *redis.Client
	Conf *config.RedisConfig
	Ctx context.Context
}
func NewRedis(conf *config.RedisConfig) (*Redis,error) {
	rdb:=&Redis{}
	rdb.WithConf(conf)
	log.Info(*conf)
	err:=rdb.ResetClient()
	if err!=nil{
		return nil,err
	}
	rdb.Ctx=context.Background()
	return rdb,nil
}
func (rdb *Redis)ResetClient() error{
	if rdb.Conf==nil{
		return errors.New("redis client configuration is null,set it first")
	}
	rdb.Client=config.NewClient(rdb.Conf)
    return nil
}

func (rdb *Redis)WithConf(conf *config.RedisConfig) {
	rdb.Conf=conf
}
func (rdb *Redis)GetClient() *redis.Client{
	return rdb.Client
}

func (rdb *Redis)SetKV(key string,value interface{}) error{
	cmd:=rdb.Client.Set(rdb.Ctx,key,value,0)
	return cmd.Err()
}
func (rdb *Redis)GetKV(key string) (string,error){
	cmd:=rdb.Client.Get(rdb.Ctx,key)
	return cmd.Result()
}
func (rdb *Redis)SetKVWithExpiration(key string,value interface{},expiration time.Duration) error{
	 cmd:=rdb.Client.Set(rdb.Ctx,key,value,expiration)
	 return cmd.Err()
}
func (rdb *Redis)SetNx(key string,value interface{}) (bool,error){
	cmd:=rdb.Client.SetNX(rdb.Ctx,key,value,0)
	return cmd.Val(),cmd.Err()
}
func (rdb *Redis)SetNxWithExpiration(key string,value interface{},expiration time.Duration) (bool,error){
	cmd:=rdb.Client.SetNX(rdb.Ctx,key,value,expiration)
	return cmd.Val(),cmd.Err()
}

func (rdb *Redis)