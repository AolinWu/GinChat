package config

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"ginchat/log"
	"time"
    "os"
	sql_log "log"
	"github.com/spf13/viper"
	"gorm.io/gorm/logger"
)

type MySQLConfig struct {
	DBName   string `yaml:"dbname"`
	Port string `yaml:"port"`
	User     string `yaml:"user"`
	PassWord string `yaml:"password"`
	CharSet  string `yaml:"charset"`
	Ip   string `yaml:"ip"`
}
const SQL_LOG_PATH="sql_log.txt"

func NewSQLLogger() (logger.Interface,error){
	f,err:=os.OpenFile(SQL_LOG_PATH,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0666)
	sql_logger:=logger.New(
		sql_log.New(f,"INFO",sql_log.Llongfile),
		logger.Config{
			SlowThreshold:  time.Second, //慢SQL查询的阈值
			LogLevel:       logger.Info,  //级别
			Colorful:       true,        //彩色
		},
	)
	return sql_logger,err
}
func (sqlconfig MySQLConfig)Dsn() string{
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
	sqlconfig.User,
	sqlconfig.PassWord,
	sqlconfig.Ip,
	sqlconfig.Port,
	sqlconfig.DBName,
	sqlconfig.CharSet)
}
type RedisConfig struct{
    Addr string `yaml:"addr"`
	Password string `yaml:"password"`
	DB int `yaml:"db"`
}
func NewClient(r *RedisConfig) *redis.Client{
	rdb := redis.NewClient(&redis.Options{
        Addr:     r.Addr,
        Password: r.Password, 
        DB:       r.DB,  
		DialTimeout: 30*time.Second,
    })
	return rdb
}
type Config struct{
	SqlServer MySQLConfig `yaml:"sqlsever"`
	RedisServer RedisConfig `yaml:"redisserver"`
}
func InitConfig(config_path string,config_name string) *Config {
	defer viper.Reset()
	config:=&Config{}
	viper.SetConfigName(config_name)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(config_path)
	if err:=viper.ReadInConfig();err!=nil{
		panic(err)
	}
	if err:=viper.Unmarshal(config);err!=nil{
		panic(err)
		
	}
	log.Info(config)
	return config
}

func (config *Config)SqlConfig() MySQLConfig{
	 return config.SqlServer
}
func (config *Config)RedisConfig() RedisConfig{
	return config.RedisServer
}

