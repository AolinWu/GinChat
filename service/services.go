package service

import (
	"fmt"
	"ginchat/log"
	"ginchat/models"
	"ginchat/utils"
	"io"
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)
const config_path="config"
var db *utils.DB=nil

func init(){
     db=utils.NewDB(config_path,utils.CONFIG_NAME)
	 
}
//GetIndex
//@Tags 首页
//@Success 200 {string} Helloworld
//@Router /index [get]
func GetIndex(c *gin.Context){
	c.JSON(200,gin.H{
		"message":"pong",
	})

   
}
// GetUserList
//@Tags 首页
//@Success 200 {string} json{"code","message"}
//@Router /user/GetUserList [get]
func GetUserList(c *gin.Context){
	users:=db.FindAllUsers()
	c.JSON(200,gin.H{
		"message":users,
	})
}
//CreateUser
//@Tags 首页
//@Summary 新增用户
//@Success 200 {string} json{"code","message"}
//@param name query string false "username"
//@param password query string false "userpasswd"
//@Router /user/CreateUser [post]
func CreateUser(c *gin.Context){
	name:=c.Query("name")
	password:=c.Query("password")
	salt:=fmt.Sprintf("%08d",rand.Int31())
    crypto_password:=utils.Encode(password,salt)
	user:=models.NewUserWithCurrentTime(name,crypto_password,salt)
	user.ClientIp=c.ClientIP()
	db.InsertUser(user)
	c.JSON(200,gin.H{
		"messages":"success",
	})
}
//DeleteUser
//@Tags 首页
//@Summary 删除用户
//@param name query string false "username"
//@param password query string false "userpasswd"
//@Success 200 {string} json{"code","message"}
//@Router /user/DeleteUser [post]
func DeleteUser(c *gin.Context) {
     name:=c.Query("name")
	 password:=c.Query("password")
	 users:=db.FindUsersByNameAndPasswd(name,password)
	 var message string
	 if len(users)==0{
		message=fmt.Sprintf("do not find user with name:%s and password:%s",name,password)
	 }else{
	    db.DeleteUsers(users)
	 }
	 c.JSON(http.StatusOK,gin.H{
       "message":message,
	 })
}
//QueryUsers
//@Tags 首页
//@Summary 通过用户名查询用户
//@param name query string false "username"
//@Success 200 {string} json{"code","message"}
//@Router /user/QueryUsers [get]
func QueryUsers(c *gin.Context){
	  username:=c.Query("name")
	  users:=db.FindUsersByName(username)
	  message:=fmt.Sprintf("%v",users)
	  c.JSON(
		http.StatusOK,
		gin.H{
         "message":message,
		})
}

//QueryUsersWithPasswd
//@Tags 首页
//@Summary 通过用户名和密码查询用户
//@param name query string false "username"
//@param password query string false "userpasswd"
//@Success 200 {string} json{"code","message"}
//@Router /user/QueryUsersWithPasswd [get]
func QueryUsersWithPasswd(c *gin.Context){
	username:=c.Query("name")
	passwd:=c.Query("password")
	users:=db.FindUsersByNameAndPasswd(username,passwd)
	message:=fmt.Sprintf("%v",users)
	c.JSON(
		http.StatusOK,
		gin.H{
			"message":message,
		},
	)


}
func SendVideo(c *gin.Context) {
	file:="C:\\嗨格式录屏文件\\20240501-234838.mp4"
	f,err:=os.Open(file)
	if err!=nil{
		panic(err.Error())
	}
	info,err:=os.Stat(file)
	if err!=nil{
		panic(err.Error())
	}
	size:=info.Size()
	buf:=make([]byte,size+1024)
	log.Info(size)
	read_size,err:=f.Read(buf)
	log.Info(read_size)
	if err!=nil{
		panic(err.Error())
	}
	c.Header("Content-Type", "video/mp4")
	c.Header("Cache-Control", "cache")
    c.Header("Connection", "keep-alive")
	cur:=0
	write_batch_size:=10
	c.Stream(func(w io.Writer) bool{
		var end_size=cur+write_batch_size
        if cur<read_size{
           if cur+write_batch_size>read_size{
			  end_size=read_size
		   }
		   w.Write(buf[cur:end_size])
		   c.Writer.Flush()
		   log.Info(end_size)
		   //c.Data(http.StatusOK,"video/mp4",buf[cur:end_size])
		   cur+=write_batch_size
           return true
		}
		return false
	})
	
   
}
