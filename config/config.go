package config

import (
	"fmt"
	"ginchat/log"

	"github.com/spf13/viper"
)

type MySQLConfig struct {
	DBName   string `yaml:"dbname"`
	Port string `yaml:"port"`
	User     string `yaml:"user"`
	PassWord string `yaml:"password"`
	CharSet  string `yaml:"charset"`
	Ip   string `yaml:"ip"`
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
type Config struct{
	SqlServer MySQLConfig `yaml:"sqlsever"`
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

