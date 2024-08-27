package routers

import (
	"github.com/gin-gonic/gin"
	c "my-ecom.com/api/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping/:name", c.IUserController().Pong)
	}

	return r
}
