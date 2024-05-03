package main

import (
	"fmt"
	"ginchat/log"
	"ginchat/models"
	"ginchat/utils"
	"testing"
	"time"
	"gorm.io/gorm"
)
var db *utils.DB=nil
func init(){
	config_path:="..\\config"
	config_name:="app"
    db=utils.NewDB(config_path,config_name)
}
func TestFind(t *testing.T){
	users:=db.FindAllUsers()
	log.Info(fmt.Sprintf("%v",users))
	if len(users)!=6{
		t.Errorf("users count should be 6,but find %d users",len(users))
	}
}
func TestInsert(t *testing.T){
	user:=&models.User{
	 	Name:"wuaolin",
        PassWord: "123",
		Phone: "17770261951",
		Email: "2015400697@qq.com",
		LoginTime: time.Now(),
		LoginOutTime: time.Now().Add(time.Hour*3),
		HeartbeatTime: time.Now(),
	}
	db.InsertUser(user)
	filter_user:=&models.User{
		Name:"wuaolin",
	}
	users,err:=db.FilterUsers(filter_user,1)
	if err!=nil{
		t.Errorf(err.Error())
	}
	log.Info(fmt.Sprintf("%v",users))
	if len(users)==0{
		t.Errorf("Users insert count should greater than 0,but find zero users")
	}
}
func TestTransaction(t *testing.T){
	 original_length:=db.UserCount()
     var op utils.TransactionOperation=func(db *utils.DB, tx *gorm.DB) error {
		user:=&models.User{
			Name:"wuaolin",
		   PassWord: "123",
		   Phone: "17770261951",
		   Email: "2015400697@qq.com",
		   LoginTime: time.Now(),
		   LoginOutTime: time.Now().Add(time.Hour*3),
		   HeartbeatTime: time.Now(),
	   }
	   tx.AutoMigrate(&models.User{})
	   tx.Create(user)
	   panic("Error happened")
	   return nil
	 }
	 defer func(){
		new_len:=db.UserCount()
        if r:=recover();r!=nil{
			if new_len!=original_length{
				t.Errorf("transaction rolback failed")
			}else{
				log.Info("transaction rollback")
			}
		}else{
			if new_len==original_length{
               t.Errorf("insert failed")
			}else{
				log.Info("transaction success")
			}
		}
	 }()
	 db.BeginTransaction(op)
	 
		 
}

