package main
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ginchat/models"
	 
)
func main(){
	 dsn:="root:123@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
     db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	 if err!=nil{
		panic("failed to open db")
	 }
	 db.AutoMigrate(&models.UserBasic{})
	user:=models.UserBasic{
    Name: "wuaolin",
	PassWord: "123",
	Phone: "17770261951",
	Email: "aolinwu@zju.edu.cn",
	}
	result:=db.Create(&user)
	if result.Error!=nil{
		 panic(result.Error.Error())
	}

}

