package router
import (
	"github.com/gin-gonic/gin"
   "ginchat/service"
   "ginchat/docs"
   swaggerfiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"
)
func Router() *gin.Engine {
   r :=gin.Default()
   docs.SwaggerInfo.BasePath = ""
   r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
   r.GET("/index",service.GetIndex)
   r.GET("/user/GetUserList",service.GetUserList)
   r.POST("user/CreateUser",service.CreateUser)
   r.POST("user/DeleteUser",service.DeleteUser)
   r.GET("/video",service.SendVideo)
   r.GET("/user/QueryUsers",service.QueryUsers)
   r.GET("/user/QueryUsersWithPasswd",service.QueryUsersWithPasswd)
   return r
}