package router
import (
	"github.com/gin-gonic/gin"
   "ginchat/service"
)
func Router() *gin.Engine {
   r :=gin.Default()
   r.GET("/index",service.GetIndex)
   return r
}