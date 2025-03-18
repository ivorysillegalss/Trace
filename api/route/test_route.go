package route

import (
	"github.com/gin-gonic/gin"
	"go-quickstart/bootstrap"
)

func RegisterTestRouter(router *gin.RouterGroup, c *bootstrap.Controllers) {
	tc := c.TestController
	testGroup := router.Group("/test", tc.HelloPost)
	testGroup.GET("/hello/:param", tc.HelloGet)
}
