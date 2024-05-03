package main

import (
	"fmt"
	"ginchat/config"
	"ginchat/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)
var conf *config.RedisConfig=nil
func init(){
	config_path:="../config"
	config_name:="app"
    tmp:=config.InitConfig(config_path,config_name).RedisConfig()
	fmt.Println(tmp)
	conf=&tmp
}
func TestRedis(t *testing.T){
	var err error=nil
	var redis *utils.Redis
	var value string
	assert.NotEqual(t,conf,nil,"conf instance is null")
	redis,err=utils.NewRedis(conf)
	if err!=nil{
		t.Errorf(err.Error())
	}
	assert.NotEqual(t,redis,nil,"redis instance is null")
	err=redis.SetKV("ss","qq")
	if err!=nil{
		t.Errorf(err.Error())
	}
	value,err=redis.GetKV("ss")
	assert.Equal(t,err,nil,"error happening while query key")
	assert.Equal(t,value,"qq",value)
}