package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"ginchat/config"
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
type DB struct{
	db *gorm.DB
	conf *config.MySQLConfig
	tx_option *sql.TxOptions
}
const MAX_RECORD_CAP=200
const CONFIG_NAME="app"
func NewDB(config_path string,config_name string) *DB{
	defer func(){
		if r:=recover();r!=nil{
			panic(r)
		}
	  }()
	conf:=config.InitConfig(config_path,config_name).SqlConfig()
    return NewDBWithConf(&conf)
}
func NewDBWithConf(conf *config.MySQLConfig) *DB{
	dsn:=conf.Dsn()
	logger,err:=config.NewSQLLogger()
	if err!=nil{
		panic(fmt.Sprintf("failed to open log file %s",config.SQL_LOG_PATH))
	}
	db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{Logger: logger})
	if err!=nil{
		panic("failed to open db")
	}
	return &DB{
		db: db,
		conf: conf,
		tx_option: &sql.TxOptions{
			Isolation: sql.LevelRepeatableRead,
			ReadOnly: false,
		},
	}
}
func (db *DB)WithConfig(conf *config.MySQLConfig){
	db.conf=conf
}
func (db *DB)WithTransConfig(conf *sql.TxOptions){
	 db.tx_option=conf
}
func (db *DB)ResetDB() error{
	if db.conf==nil{
		return errors.New("initialize db conf first")
	}
	dsn:=db.conf.Dsn()
	init_db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		return err
	}
	db.db=init_db
    return nil
}
func (db *DB)InsertUser(insert_user *models.User){
	 db.db.AutoMigrate(&models.User{})
	 db.db.Create(insert_user)
}
func (db *DB)FindAllUsers() []models.User{
	users:=make([]models.User,0)
	db.db.Find(&users)
	return users
}
func (db *DB)FilterUsers(filter_user *models.User,args ...interface{}) ([]models.User,error){
	var cap int=MAX_RECORD_CAP
	ok:=false
	if len(args)>0{
        cap,ok=args[0].(int)
		if !ok{
          return []models.User{},errors.New("parameter error in args,should pass capacity in first args")
		}
	}
	users:=make([]models.User,0,cap)
	db.db.Where(filter_user).Find(&users)
	return users,nil
}

func (db *DB)DeleteUser(delete_user *models.User){
	db.db.Unscoped().Delete(delete_user)
}
func (db *DB)DeleteUsers(delete_users []models.User){
	db.db.Unscoped().Delete(&delete_users)
}
func (db *DB)UpdateUser(original_user models.User,updated_user models.User){
     db.db.Model(&original_user).Updates(&updated_user)
}
type TransactionOperation func(db *DB,tx *gorm.DB) error
func (db *DB)BeginTransaction(op TransactionOperation){
     tx:=db.db.Begin()
     err:=op(db,tx)
	 if err!=nil{
		tx.Rollback()
	 }else{
		tx.Commit()
	 }
}
func (db *DB)UserCount() int{
	return len(db.FindAllUsers())
}

func (db *DB)FindUsersByName(name string) []models.User{
	  user:=models.NewUser()
	  user.Name=name
	  users,err:=db.FilterUsers(user)
	  if err!=nil{
		panic("Error Happened")
	  }
	  return users
}

func (db *DB)FindUsersByNameAndPasswd(name string,passwd string) []models.User{
	  users:=db.FindUsersByName(name)
	  ans:=make([]models.User,0,1)
	  for _,user:=range users{
		  if Verify(passwd,user.Salt,user.PassWord){
             ans=append(ans, user)
		  }
	  }
	  return ans
}